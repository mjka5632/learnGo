package main

import "fmt"

type Rectangle struct {
	length, width float32
}




func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area of this shape is", shapes[n].Area())
	}
}
