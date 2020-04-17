// File:service.go
// Date:2020/4/15
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iikira/Baidu-Login"
	"net/http"
	"strings"
	"xpan/internal/pcsconfig"
)

// client请求数据Model
type RequestModel struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

// service返回Model
type ResponseModel struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var (
	upgrader = websocket.Upgrader{
		Error: onError,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// service类
type Service struct {
	bc       *baidulogin.BaiduClient
	lj       *baidulogin.LoginJSON
	vcodeStr string
}

func (s *Service) init() {
	s.bc = baidulogin.NewBaiduClinet()
}

// 发送验证码
func (s *Service) sendMobileCode() *ResponseModel {
	msg := s.bc.SendCodeToUser("mobile", s.lj.Data.Token)

	return &ResponseModel{
		Code:    0,
		Message: msg,
		Success: true,
	}
}

// 登录
func (s *Service) login(model *RequestModel) *ResponseModel {
	var result = new(ResponseModel)
	user := strings.Split(model.Data, "|")

	var bduss, ptoken, stoken, codeAddr string

	s.lj = s.bc.BaiduLogin(user[0], user[1], user[2], s.vcodeStr)

	switch s.lj.ErrInfo.No {
	case "0": // 登录成功, 退出循环
		bduss = s.lj.Data.BDUSS
		ptoken = s.lj.Data.PToken
		stoken = s.lj.Data.SToken
	case "400023", "400101": // 需要验证手机或邮箱
		nlj := s.bc.VerifyCode("mobile", s.lj.Data.Token, user[2], s.lj.Data.U)

		if nlj.ErrInfo.No != "0" {
			// 发送给前端错误数据
			panic("校验用户手机验证码异常")
		}
		// 登录成功
		bduss = s.lj.Data.BDUSS
		ptoken = s.lj.Data.PToken
		stoken = s.lj.Data.SToken
	case "500001", "500002": // 验证码

		s.vcodeStr = s.lj.Data.CodeString
		if s.vcodeStr == "" {
			panic("获取验证码异常")
		}

		// 图片验证码地址
		codeAddr = "https://wappass.baidu.com/cgi-bin/genimage?" + s.vcodeStr

	default:
		panic("登录异常")
	}

	if bduss != "" && ptoken != "" && stoken != "" {
		_, err := pcsconfig.Config.SetupUserByBDUSS(bduss, ptoken, stoken)
		if err != nil {
			panic(err)
		}
		result = &ResponseModel{
			Code:    0,
			Message: "登录成功",
			Success: true,
		}
	} else {

		result = &ResponseModel{
			Code:    1,
			Message: codeAddr,
			Success: false,
		}

	}

	return result
}

//
func (s *Service) logout(model *RequestModel) *ResponseModel {
	return nil
}

// 获取账号列表
func (s *Service) loglist(model *RequestModel) *ResponseModel {
	return nil
}

// 获取用户信息
func (s *Service) who(model *RequestModel) *ResponseModel {

	return nil
}

// 接收客户端数据
func (s *Service) received(ws *websocket.Conn) {
	// 读取客户端数据
	mt, data, err := ws.ReadMessage()
	if err != nil {
		panic(err)
	}

	// 解析返回的字节数组
	model := new(RequestModel)
	if err := json.Unmarshal(data, &model); err != nil {
		fmt.Println(err)
	}

	// 根据解析到指令执行对应的操作
	result := new(ResponseModel)
	switch model.Cmd {
	case "who":
		result = s.who(model)
	case "login":
		result = s.login(model)
	case "sendMobileCode":
		result = s.sendMobileCode()

	}

	// 结果返回给前端
	bresult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	err = ws.WriteMessage(mt, bresult)

	if err != nil {
		panic(err)
	}

	s.received(ws)
}

// 服务端通用方法,解析用户请求,执行对应指令
func (s Service) Do(w http.ResponseWriter, r *http.Request) {
	// 初始化ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ws.Close()
		}
	}()
	s.received(ws)

}

// 创建新的service对象
func NewService() *Service {
	return &Service{}
}

// ws异常处理
func onError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	fmt.Println("当前异常状态:", status, ",异常原因:", reason)
}
