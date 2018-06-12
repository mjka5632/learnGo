package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	//ret++是在执行return 1语句后发生的
	return 1
}
func main() {
	//结果为2
	fmt.Println(f())
}
