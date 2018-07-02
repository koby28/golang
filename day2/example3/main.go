package main

import "fmt"

func main()  {
	type Person struct {
		name string    //名字
		sex byte    //性别
		age int    //年龄
	}

	type Student struct {
		Person     //继承
		id int    //ID
		addr string    //地址
	}

	var s = Student{Person{"mike",'m',18},1,"bj"}
	fmt.Printf("s=%+v\n",s)

	//自动推导类型
	s2 := Student{Person{"mike",'m',19},2,"sh"}
	//fmt.Println("s2=",s2)
	//%+v显示更详细的信息
	fmt.Printf("s2=%+v\n",s2)

	//指定成员初始化，没有初始化的自动赋值为0
	s3 := Student{id:1}
	fmt.Println("s3=",s3)

	s4 := Student{ Person : Person{name:"mike"},id:1}
	fmt.Printf("s4=%+v\n",s4)



}
