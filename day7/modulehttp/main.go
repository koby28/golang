package main

import (
	"html/template"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Title string
	age string
}

func main()  {
	t,err := template.ParseFiles("G:/StudyGo/src/LearGoProject/day7/modulehttp")
	if err != nil{
		fmt.Println("parse file err:",err)
		return
	}
	p := Person{"Mary","我的个人网站！","20"}
	if err := t.Execute(os.Stdout,p);err != nil{
		fmt.Println("There was an error:",err.Error())
	}
}
