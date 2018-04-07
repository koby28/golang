package main

import "fmt"

func add(a int,arg... int) int  {
    var sum int = a
	for i := 0;i<len(arg);i++  {
		sum += arg[i]
	}
	return sum
}

func main()  {
	const (
		Unknown = iota
		Female
		Male
	)
	fmt.Println(Unknown,Female,Male)
	sum := add(3,4,54,3)
	fmt.Println(sum)
	
}
