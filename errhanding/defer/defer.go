package main

import (
	"fmt"
	"os"
	"bufio"
	"learnGo/functional/fib"
)

func tryDefer() {
	//defer 不怕return panic 都会打印出来
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}
func writeFile(filename string) {
	//create newWriter开了
	// 就需要defer关掉
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//file, err := os.Create(filename)
	//err=errors.New("this is a custom error")/
	if err != nil {
		//fmt.Println("file already exists")
		//fmt.Println("Error:",err)
		//err是否属于*os.PathError，不属于panic
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {

			fmt.Printf("%s, %s,%s\n" ,
			pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	writeFile("aa1.txt")
	//tryDefer()
}
