package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		//recover仅仅在 defer中调用，出错之前进行捕获
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))

	//b:=0
	//a := 5 / b
	//fmt.Println(a)

	panic(123)
}
func main() {
	tryRecover()
}
