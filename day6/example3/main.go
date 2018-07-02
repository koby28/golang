package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(str string)  {
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1()  {
	Printer("hello")
	ch <- 888  //给管道写数据，发送
}
func person2()  {
	<- ch
	Printer("world")
}

func main()  {
    go person1()
    person2()
}
