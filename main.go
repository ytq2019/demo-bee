package main

import (
	_ "demo-bee/routers"
	_ "demo-bee/task"
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()
}
