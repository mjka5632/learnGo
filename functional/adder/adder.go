package main

import "fmt"

func add() func(int) int {
	//闭包
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+....+%d=%d\n", i, a(i))
	}
}
