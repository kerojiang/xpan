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
type cmd struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}

var (
	upgrader = websocket.Upgrader{
		Error: onError,
	}
	conn *websocket.Conn
)

// service类
type Service struct {
}

func (s Service) login(ws *websocket.Conn, msg string) {

}

func (s Service) logout(ws *websocket.Conn, msg string) {

}

// Begin
func (s Service) Do(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
}

// OnReceive
func (s Service) OnReceive() (cmd, error) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			return cmd{}, err
		}

		// 解析返回的字节数组
		clientCmd := cmd{}
		if err := json.Unmarshal(data, &clientCmd); err != nil {
			return clientCmd, nil
		}
	}
}

// Close
func (s Service) Close() bool {
	err := conn.Close()
	if err != nil {
		return false
	}
	return true
}

// 异常处理
func onError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	fmt.Println("当前异常状态:", status, ",异常原因:", error.Error())
}
