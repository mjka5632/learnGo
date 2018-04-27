package main

import "fmt"

func printSlice(a []int) {
	fmt.Printf("val= %v,len= %d,cap= %d\n", a[:], len(a), cap(a))

}
func main() {
	s := []int{2, 3, 4, 5, 6, 7}
	fmt.Println("Create silce")
	//make(类型，len(),cap())
	a := make([]int, 6, 18)
	//make(type(),len())
	b := make([]int, 16)
	fmt.Println(a[1], b)
	fmt.Println("Copying slice")
	//s的值拷贝到了a中
	copy(a, s)
	copy(b, s)
	fmt.Println(a)
	fmt.Println(b)
	printSlice(b)
	fmt.Println("Deleting elements")
	b = append(b[:2], b[3:]...)
	fmt.Println(b)
	printSlice(b)
	fmt.Println("Popping from front")
	t := b[0]
	b=b[1:]
	fmt.Println(t)
	printSlice(b)
	fmt.Println("Popping from back")
	//q:=b[len(b)-1]
	b=b[:len(b)-1]
	//fmt.Println(q)
	printSlice(b)

}
