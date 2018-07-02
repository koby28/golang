package main

import "fmt"

func xxx(arg ...interface{})  {
	
}
func main()  {
	//空接口是万能类型的，可以保存任意类型的值
	var i  interface{} = 1
	fmt.Println("i=",i)

	i = "abc"
	fmt.Println("i=",i)
    xxx("aaa")
}
