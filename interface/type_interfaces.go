package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}
type Circle struct {
	radius float32
}

func (r *Circle) Area() float32 {
	return r.radius * r.radius * math.Pi
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is : %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is : %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")

	}
}
