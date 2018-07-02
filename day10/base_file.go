package main

import (
	"os"
	"github.com/gpmgo/gopm/modules/log"
	"fmt"
)

func readFile(path string)  {
	file, err := os.Open(path)
	if err != nil{
		log.Fatal("err : %v", err)
	}
	buf := make([]byte, 1024)
	for  {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		fmt.Println(buf[:n])
	}
}

func main() {
	fmt.Println("hello golang")
	readFile("test.txt")
}
