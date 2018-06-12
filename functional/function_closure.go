package main

import "fmt"

func main() {
	var f = Adder1()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))

}

//闭包函数体内声明
func Adder1() func(int) int {
	//初始化x=0
	//在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321：
	// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	var x int
	return func(delta int) int {
		x += delta
		return x
	}

}

//闭包函数体外声明
//var g int
//go func (i int) {
//	s := 0
//	for j := 0; j < i; j++ {
//		s += j
//	}
//	g = s
//}(1000)
