package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("aaaaaaaaaaaaaaaaaaa")
	runtime.Goexit()
	fmt.Println("bbbbbbbbbbbbbbbbbbb")
}

func main() {
	go func() {
		fmt.Println("cccccccccccccc")
	}()
	test()
	fmt.Println("ddddddddddddd")
	for{}
}
