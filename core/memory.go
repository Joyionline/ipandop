package core

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	Cacheinfo cache.Cache
	err       error
)

func init() {
	Cacheinfo, err = cache.NewCache("memory", `{"interval":5}`)
	if err != nil {
		beego.Error("Error：", err)
	}
	Cacheinfo.Put("hello", "world", time.Second*1500)
	beego.Info("初始化缓存数据：", Cacheinfo)
}
