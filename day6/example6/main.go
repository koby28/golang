package main

import (
	"time"
	"fmt"
)

func main() {
	//timer := time.NewTimer(3*time.Second)   //定时器
	//
	//go func() {
	//	<- timer.C
	//	fmt.Println("子协程可以打印了，因为定时器时间到")
	//}()
	//timer.Stop()   //停止定时器
	//for{}
	timer := time.NewTimer(3*time.Second)
    timer.Reset(1*time.Second)    //重置定时器
	<- timer.C
	fmt.Println("时间到")
}
