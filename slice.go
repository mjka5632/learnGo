package main

import (
	"fmt"
)

//切片是引用类型，不能使用指针
var x = []int{2, 4, 5, 8, 9}

//切片当做参数
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s

}
func main() {
	arr3 := [5]int{1, 2, 3, 4, 5}
	//只要下标
	for i := range arr3 {
		//fmt.Println(arr3[i])
		fmt.Println(i)
	}
	//下标和值都要
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//只要值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	arr4 := [...]int{1, 2, 3, 4, 10, 8}
	maxi := -1
	maxVaue := -1
	//循环打印出最大值
	for i, v := range arr4 {
		if maxVaue < v {
			maxi, maxVaue = i, v
		}
	}
	fmt.Println(maxi, maxVaue)
	//代表四行五列
	var grid [4][5]int
	fmt.Println(grid)
	arr := []int{0, 1, 2, 3, 4}
	s := sum(arr[:])
	fmt.Println("sum %d", sum(arr[:]))
	fmt.Println("sum %d", s)
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i

	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Println("arr1 is %d", len(arr1))
	fmt.Println("slice1 is %d", len(slice1))
	fmt.Println("slice1 cap %d", cap(slice1))
	fmt.Println(slice1)
}
