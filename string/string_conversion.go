package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig = "ABC"
	var newS string

	fmt.Printf("The size of ints is %d\n", strconv.IntSize)

	an, e := strconv.Atoi(orig)
	if e != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n", newS)
}
