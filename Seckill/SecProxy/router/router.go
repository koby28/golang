package router

import (
	"github.com/astaxie/beego"
	"LearGoProject/Seckill/SecProxy/controller"
)


func init() {
	beego.Router("/seckill",&controller.SkillController{},"*:SecKill")
	beego.Router("/secinfo",&controller.SkillController{},"*:SecInfo")
}
