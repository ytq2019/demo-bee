package routers

import (
	"demo-bee/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/index", &controllers.IndexController{})
	beego.Router("/api/list/:type([\\w]+)/:page([0-9]+)", &controllers.ListController{})
	beego.Router("/api/detail", &controllers.DetailController{})
}
