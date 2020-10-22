package controllers

import (
	"demo-bee/inc"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
)

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {
	detail := inc.GetVideoDetail("/tv/PrJob07lSz8tOX.html", "dianshi")
	c.Data["json"] = detail
	c.ServeJSON()
}
