// File:main.go
// Date:2020/4/15
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	xpanService := NewService()

	// 注册静态文件位置
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("./www/"))))

	// 注册login静态页面
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./www/login.html")
		if err != nil {
			fmt.Println("启动页面异常:", err)
		}
		_ = t.Execute(writer, nil)
	})
	// 注册websocket方法
	http.HandleFunc("/ws", xpanService.Do)

	fmt.Println("请打开浏览器,访问 http://localhost:8989")

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("启动服务端失败:", err)
	}

}
