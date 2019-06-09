package controller

import (
	"net/http"

	"github.com/astaxie/beego"
)

func Cache2go(rsp http.ResponseWriter, req *http.Request) {
	beego.Info("Request：", req.Host, req.RemoteAddr)
}
