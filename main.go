// File:main.go
// Date:2020/4/15
package main

import (
	"fmt"
	"net/http"
)

func main() {

	xpanWS := Service{}
	http.HandleFunc("/ws", xpanWS.Do)
	err := http.ListenAndServe("127.0.0.1:8989", nil)
	if err != nil {
		fmt.Println("启动服务端失败:", err)
	}
	fmt.Println("启动服务端成功")
}
