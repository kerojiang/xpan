// File:service.go
// Date:2020/4/15
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// 数据Model
type CmdModel struct {
	Key string `json:"key"`
	Cmd string `json:"cmd"`
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
}

// 登录
func (s Service) login(msg string) string {
	return ""
}

//
func (s Service) logout(msg string) string {
	return ""
}

// 获取账号列表
func (s Service) loglist(msg string) string {
	return ""
}

// 获取用户信息
func (s Service) who(msg string) string {

	return "123"
}

// 接收客户端数据
func (s Service) received(ws *websocket.Conn) {
	// 读取客户端数据
	mt, data, err := ws.ReadMessage()
	if err != nil {
		panic(err)
	}

	// 解析返回的字节数组
	clientCmd := CmdModel{}
	if err := json.Unmarshal(data, &clientCmd); err != nil {
		fmt.Println(err)
	}

	// 根据解析到指令执行对应的操作
	var result string
	switch clientCmd.Key {
	case "who":
		result = s.who(clientCmd.Cmd)
	case "login":
		result = s.login(clientCmd.Cmd)

	}

	// 结果返回给前端
	err = ws.WriteMessage(mt, []byte(result))

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

// 异常处理
func onError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	fmt.Println("当前异常状态:", status, ",异常原因:", reason)
}
