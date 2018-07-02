package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	//创建一个map
	m := make(map[string]interface{},4)
	m["company"] = "51CTO"
	m["subject"] = []string{"go","java","C++","python"}
	m["isok"] = true
	m["price"] = 666.666
	//编码成json
	result,err := json.MarshalIndent(m,""," ")
	if err != nil{
		fmt.Println("MarshalIndent error",err)
		return
	}
	fmt.Println("result=",string(result))
}
