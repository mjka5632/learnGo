package main

import "fmt"

func f() (ret int) {
	//因为匿名函数没有名称。
	// 花括号 {} 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。
	defer func() {
		ret++
	}()
	//ret++是在执行return 1语句后发生的
	return 1
}
func main() {
	//结果为2
	//fmt.Println(f())
	//匿名函数()表示赋值
	//func(x,y int) int{return x+y}(3,4)
	//f2 := func(x, y int) int { return x + y }
	//f2(3, 4)
	//fmt.Print(f2(3, 4))
	f2()
}
func f2() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}