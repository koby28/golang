package router

import (
	"github.com/astaxie/beego"
	"LearGoProject/day8/beego_demo/controller/IndexController"
)

func init()  {
	beego.Router("/",&IndexController.IndexController{},"*:Index")
	beego.Router("/index",&IndexController.IndexController{},"*:Index")
}
