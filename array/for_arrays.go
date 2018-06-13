package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	//for的生成方式
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d \n", i, arr1[i])
	}
	//for-range的生成方式
	for i,_:=range arr1{
		fmt.Printf("Array at index %d is %d \n", i, arr1[i])

	}
}
