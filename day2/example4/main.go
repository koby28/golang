package main

import "fmt"

type Humaner interface {
	sayhi()
} 

type Personer interface {
	Humaner
	sing(lrc string)
}

type Student struct {
	name string
	id int
}

//Student实现sayhi()
func (tmp *Student) sayhi()  {
	fmt.Printf("Student [%s,%d] sayhi\n",tmp.name,tmp.id )
}

func (tmp *Student) sing(lrc string)  {
	fmt.Printf("Student在唱歌：",lrc)
}

func main()  {
	//定义一个接口类型的变量
	var i Personer
	s := &Student{"mike",666}
	i = s
	i.sayhi()
	i.sing("唱起了")
}