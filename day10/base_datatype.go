package main

import "fmt"

//定义常量
const (
	a int = 1
	b string = "test"
	flag bool = false
	c        = "33" //类型可选项
)

func testArray()  {
	//一维数组
	var arr [2]int //这样定义不会报未使用错误
	arr[0] = 1
	arr[1] = 1

	var arr2 [2]int = [2]int{1,2} //简写var arr2 = [2]int{1,2}
	fmt.Printf("arr2 = %V\n",arr2)
	var arr3  = [...]int{1,2,3,}
	fmt.Printf("arr3 = %v\n",arr3)

	//二维数组
	var arr4 [2][2]int = [2][2]int{[2]int{1,2},[2]int{2,3}}
	fmt.Printf("arr4 = %v\n",arr4)
	var arr5 = [2][2]int{[...]int{1,2},[...]int{3,4}}
	fmt.Printf("arr5 = %d\n",arr5)
	var arr6 = [2][2]int{{2,3},{4,5}}
	fmt.Printf("arr6 = %d\n",arr6)
}
/**
:= 的含义：同时进行变量的声明及初始化工作.
数组和slice的区别
a.声明数组时，方括号内写明了数组的长度或者...,声明slice时候，方括号内为空
b.作为函数参数时，数组传递的是数组的副本，而slice传递的是指针。
**/
func testSlice() {
	//var arr1 []int	  //创建slice
	//arr2 := make([]int, 2) //创建slice
	arr := []int{1, 2}

	//对于数组切片
	var tmp = [3]int{1, 2, 3}
	arr = tmp[0:2]

	//slice append
	arr = append(arr, 1, 2, 3)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d = %d\n",i, arr[i])
	}
}
func testMap() {
	var m map[string]int                               //创建map
	m = make(map[string]int)                           //创建map
	m = map[string]int{"one": 1, "two": 2, "three": 3} //创建有初始化元素的map
	m["one"] = 5

	//遍历
	for key, val := range m {
		fmt.Println(key, val)
	}
	//判断key是否存在，v 为value ,ok 为标识-ture|false
	v, ok := m["e"]
	fmt.Print(v, ok)

	//删除键
	delete(m, "one")

	fmt.Printf("m = %v",m)
}

func main()  {
	//testArray()
	testSlice()
	//testMap()
}