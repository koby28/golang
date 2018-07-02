package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	list := os.Args //获取命令行参数

	if len(list) != 3 {
		fmt.Printf("usage: xxx srcFile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName{
		fmt.Println("源文件和目的文件不能相同")
		return
	}

	//只读方式打开原文件
	sF,err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1=",err1)
		return
	}

	//新建目的文件
	dF,err2 := os.Create(dstFileName)
	if err2 != nil{
		fmt.Println("err2=",err2)
		return
	}
	//操作完毕，需要关闭文件
	defer sF.Close()
	defer dF.Close()
	//核心处理，从原文件读取内容，往目的地写，读多少写多少

	buf := make([]byte,4*1024)
	for{
		n,err := sF.Read(buf)
		if err != nil{
			fmt.Println("err=",err)
			if err == io.EOF {
				break
			}
		}
		dF.Write(buf[:n])
	}

}
