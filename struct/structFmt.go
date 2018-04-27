package main

import "fmt"
//struct转换
type number struct {
	f float32
}
type nr number

func main() {

	a := number{5.0}
	b := nr{5.0}
	var c = number(b)
	//var i float32=b		float32定义都不可以啦，已经没有关系啦
	//var i=float32(b)
	//var c number=b
	fmt.Println(a,b,c)
}
