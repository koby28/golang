package main

import (
	"fmt"
	errors2 "errors"
)

/**
基本语法 “=”与“:=”的区别
 */
var td  = "dd"

const version  = "1.0" //常量定义，不可修改
func main()  {
	var i  = `ddd`
	j := 0 //只能用到函数内部
	fmt.Printf("%sImy%d\n",i,j)
	td = "dd0"
	fmt.Print(td)
	var a, b, c = 1, 2, 3
	fmt.Println(a,b,c)
	var d,_ = 1.22, 33    //特殊变量
	fmt.Println(d)
	error()
}

func error()  {
	err := errors2.New("测试肺胀")
	if err != nil {
		fmt.Println(err.Error())
	}
}
