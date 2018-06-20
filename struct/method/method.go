package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

//(tn *TwoInts)值接收者
func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
func (tn *TwoInts) String() string {
	return "("+ strconv.Itoa(tn.a)+"/"+strconv.Itoa(tn.b)+")";
}
func main() {
	t1 := new(TwoInts)
	t1.a = 12
	t1.b = 10
	fmt.Printf("The sum is: %d\n", t1.AddThem())
	fmt.Printf("Add them to the param: %d\n", t1.AddToParam(20))
	//String改变样式
	fmt.Println("two1 is:", t1)
	fmt.Printf("two1 is: %T\n", t1)

	t2 := TwoInts{3, 4}
	fmt.Printf("The sum is : %d\n", t2.AddThem())
}
