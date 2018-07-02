package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaaaa")
}

func testb(x int)  {
	//fmt.Println("bbbbbbbbbbbbb")
	defer func() {
		if err := recover();err != nil  {
			fmt.Println(recover())
		}
	}()
	var a [10]int
	a[x] = 100
	fmt.Printf("a[%x]=%a\n",x,a[x])
	//panic("this is a panic test")
}

func testc()  {
	fmt.Println("cccccc")
}


func main()  {
	testa()
	testb(9)
	testc()
}
