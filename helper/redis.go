/*
 * @Time : 2020/2/22
 * @Author : tianqi.yu
 * @Software : GoLand
 * @Description :
 */
package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var adapter *cache.Cache = nil

func RedisAdater() *cache.Cache {
	if adapter == nil {
		bm, err := cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s","dbNum":"0","password":%s}`, beego.AppConfig.String("redishost"), beego.AppConfig.String("redispwd")))
		if err != nil {
			panic(err)
		}
		adapter = &bm
	} else {
		return adapter
	}

	return adapter
}
