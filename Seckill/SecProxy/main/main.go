package main

import (
	"github.com/astaxie/beego"
	_ "LearGoProject/Seckill/SecProxy/router"
)

func main() {
	err := initConfig()
	if err != nil{
		panic(err)
		return
	}

	err = initSet()
	if err != nil{
		panic(err)
		return
	}
	beego.Run()
}
