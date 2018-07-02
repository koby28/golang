package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	result := `
{
		"company": "51CTO",
			"isok": true,
			"price": 666.666,
			"subject": [
		"go",
			"java",
			"C++",
			"python"
		]
	}`
	//创建一个map
	m := make(map[string]interface{},4)
	err := json.Unmarshal([]byte(result),&m)
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	fmt.Printf("m=%+v\n",m)
}
