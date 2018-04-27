package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	//使用空接口
	var val Any
	fmt.Printf("value:%v\n", val)
	val = str
	fmt.Printf("value:%v\n", val)

	p1 := new(Person)
	p1.name = "Rob Pike"
	p1.age = 55
	val = p1
	fmt.Printf("value:%v\n", val)

	switch t := val.(type) {
	case *Person:
		fmt.Printf("Type %T\n",t)
	case string:
		fmt.Printf("Type %T\n",t)
	default:
		fmt.Printf("无此类型",t)
	}
}
