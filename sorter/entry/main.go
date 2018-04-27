package main

import (
	"studygolang/sorter"
	"fmt"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data)
	//进行排序
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is :%v\n", a)
}

func strings() {
	data := []string{"One", "Two", "Three", "Five", "Six"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}
func main() {
	ints()
	strings()
}
