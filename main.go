// File:main.go
// Date:2020/4/15
package main

import (
	"net/http"
)

func main() {

	xpanWS := Service{}
	http.HandleFunc("/ws", xpanWS.Do)
	http.ListenAndServe("127.0.0.1:8989", nil)
}
