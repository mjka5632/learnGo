package main

import (
	"learnGo/functional/fib"
	"fmt"
)

func main() {

	result := 0
	for i := 0; i <= 4; i++ {
		result = fib.Fibonacci1(i)
		fmt.Printf("fibonacci(%d)is : %d\n", i, result)
	}
}
