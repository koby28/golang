package main

import "fmt"

func main() {
	ch := make(chan int) //创建一个无缓存区的channel
	go func() {
		for i := 0;i<5;i++ {
			ch <- i // 往管道里写数据
		}
        close(ch)
	}()
	for{
		//如果OK为true，说明管道没有关闭
		if num,ok := <- ch;ok==true {
			fmt.Println("num=",num)
		}else {
			break
		}
	}
}
