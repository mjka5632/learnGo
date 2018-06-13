package main

import "fmt"

func main() {
	//map放入切片中
	items := make([]map[int]int, 5)
	for i := range items {
		//map
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items:%v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B:Value of items: %v\n", items2)
}
