package main

import (
	"fmt"
	"runtime"
)

func main() {

	for i :=0;i<2;i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
	//匿名函数
	go func() {
		for i := 0;i<5 ;i++  {
			fmt.Println("go")
		}
	}()
}
