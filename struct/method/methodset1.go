package main

import "fmt"


func main() {
	//值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len())

	//指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len())
}
