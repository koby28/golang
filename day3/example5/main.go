package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func WriteFile(path string) {
    //打开文件，新建文件
    f,err := os.Create(path)
	if err != nil {
        fmt.Println("err=",err)
        return
	}
	//使用延时关闭文件
	defer f.Close()

	var buf string
	for i:=0;i<10 ;i++  {
		buf = fmt.Sprintf("i=%d\n",i)
		fmt.Println("buf=",buf)

		n,err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err=",err)
		}
		fmt.Println("n=",n)
	}
}

//读取文件
func ReadFile(path string)  {
	//打开文件
	f,err := os.Open(path)
	if err != nil {
		fmt.Println("err=",err)
		return
	}
	//关闭文件
	defer f.Close()

	buf := make([]byte,1024*2)   //2K大小
	//n代表从文件读取内容长度
	n,err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println("err1=",err1)
		return
	}
	fmt.Println("buf=",string(buf[:n]))

}

//根据读取行
func LineFile(path string)  {
	//打开文件
	f,err := os.Open(path)
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	//关闭文件
	defer  f.Close()
	//新建一个缓冲区，把内容放入缓冲区中
	r := bufio.NewReader(f)

	for{
		buf,err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已结束
				break
			}
			fmt.Println("err=",err)
		}
		fmt.Printf("buf=#%s#\n",string(buf))
	}

}

func main() {
    path := "./test.txt"
    //WriteFile(path)
    //ReadFile(path)
    LineFile(path)
}
