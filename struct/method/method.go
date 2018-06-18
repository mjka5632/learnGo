package main

import "fmt"

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
func main() {
	t1 := new(TwoInts)
	t1.a = 12
	t1.b = 10
	fmt.Printf("The sum is: %d\n", t1.AddThem())
	fmt.Printf("Add them to the param: %d\n", t1.AddToParam(20))

	t2 := TwoInts{3, 4}
	fmt.Printf("The sum is : %d\n", t2.AddThem())
}
