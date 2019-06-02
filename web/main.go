package main

import (
	"net/http"
	"tools"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func handle(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte(tools.GetRemoteIP(rq)))
}
