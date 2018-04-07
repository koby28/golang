package main

import "fmt"
//数组作为函数，它是值传递
//实参数组的每个元素给形参数组拷贝一份
//形参数组是实参数组的复制品
func main()  {
	a := [6]int {3,2,1,5,4,6}

	//modify(a)
	modify(&a)
	fmt.Println("main a=",a)


}
/*
func modify(a [6]int)  {
	a[0] = 66666
	fmt.Println("modify a=" , a)
}*/
func modify(p *[6]int)  {
	(*p)[0] = 66666
	fmt.Println("modify a=" , *p)

}




