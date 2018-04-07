package main

import (
	"math/rand"
	"time"
	"fmt"
)

//冒泡排序
func main()  {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	var a [10]int
	n := len(a)
	for i:=0;i<n ;i++  {
		a[i] = rand.Intn(100) //100以内的随机数
		fmt.Printf("%d",a[i])
		fmt.Printf("\n")
	}
	fmt.Println("-----------------------------")
	//冒泡业务逻辑
	for i :=0;i<n-1 ;i++  {
		for j :=0;j<n-1-i;j++ {
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
	for i:=0;i<n ;i++  {
		fmt.Printf("%d",a[i])
		fmt.Printf("\n")
	}

}

