package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	ch := make(chan string )
	go func() {
		defer fmt.Println("子协程调用完毕！")
		for i := 0;i<2 ;i++  {
			fmt.Println("子协程 i=",i)
			time.Sleep(time.Second)
		}
        ch <- "我是子协程，要工作完毕！"
	}()
	str := <-ch
	fmt.Println("str=",str)
}
