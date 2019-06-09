package main

import (
	"fmt"
	"ipandop/controller"
	"net/http"
	"tools"
)

var (
	address string
	port    string = ":80"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/cache2go", controller.Cache2go)
	fmt.Println("http server Running on http://", port)
	http.ListenAndServe(port, nil)
}

// 获取IP
func handle(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte(tools.GetRemoteIP(rq)))
}
