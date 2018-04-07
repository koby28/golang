package main

import (
	"LearGoProject/day2/test"
	"fmt"
)
//使用不同包的方法或结构体
func main()  {
	test.MyFunc()
	var s  test.Stu
	var ss  = test.Stu{18}
	s.Id = 555
	fmt.Println("s=",ss)
}
