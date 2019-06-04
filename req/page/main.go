package main

import (
	"sync"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/httplib"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(i int) {
			beego.Info("访问第", i, "次")
			req := httplib.Get("")
			_, err := req.String()
			if err != nil {
				beego.Error("Error:", err)
			}
			// beego.Info(s)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
