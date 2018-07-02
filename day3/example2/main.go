package main

import (
	"strconv"
	"fmt"
)

func main()  {
	slice := make([]byte,0,1024)
	slice = strconv.AppendBool(slice,true)
	//第二个参数要追加的数，第三个为指定的10进制方式追加
	slice = strconv.AppendInt(slice,1234,10)
	slice = strconv.AppendQuote(slice,"abcd")
	fmt.Println("slice=",string(slice))

	//其他类型转换字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)

    str = strconv.Itoa(666)
    fmt.Println("str=",str)

    //字符串转化为其他类型
    var flag bool
    var err error
    flag,err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ",flag)
	}else{
		fmt.Println("err=",err)
	}

	//将字符串转化成整型
	a,_ := strconv.Atoi("10")
	fmt.Println(a)


	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[8:]
	fmt.Println(s1)
	s2 := data[:5]
	fmt.Println(s2)
	copy(s2, s1)
	fmt.Println(data)


}
