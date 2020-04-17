// File:service.go
// Date:2020/4/15
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
	"net/http"
	"strings"
	"xpan/internal/pcscommand"
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
	c *cli.Context
)

// service类
type Service struct {
}

// 登录
func (s *Service) login(model *RequestModel) *ResponseModel {

	user := strings.Split(model.Data, "|")

	bduss, ptoken, stoken, err := pcscommand.RunLogin(c.String(user[0]), c.String(user[1]))
	if err != nil {
		panic(err)
	}
	_, err = pcsconfig.Config.SetupUserByBDUSS(bduss, ptoken, stoken)
	if err != nil {
		panic(err)
	}

	result := &ResponseModel{
		Code:    0,
		Message: "登录成功",
		Success: true,
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
	// 初始化
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
