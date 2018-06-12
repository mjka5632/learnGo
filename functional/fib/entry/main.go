package main

import (
	"time"
	"learnGo/functional/fib"
	"fmt"
)

func main() {
	var result uint64 = 0
	start:=time.Now()
	for i:=0;i<fib.LIM;i++{
		result=fib.FibonacciMe(i)
		fmt.Printf("fibonacci_Me(%d) is : %d\n",i,result)
	}
	end:=time.Now()
	delta:=end.Sub(start)
	fmt.Printf("LongCalculation took this amount of time:%s\n",delta)
}
