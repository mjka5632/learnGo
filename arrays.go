package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [...]int{1, 2, 3, 4}
	arr3 := [5]int{1, 2, 3, 4, 5}
	//代表四行五列
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
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

}
