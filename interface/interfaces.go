package main

import "fmt"



func main() {
	sq1 := new(Square)
	sq1.side = 5
	//declare 1
	//var areaIntf Shaper
	//areaIntf=sq1
	//declare 2
	//areaIntf:=Shaper(sq1)
	//declare3
	//area := sq1.Area()
	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	//fmt.Printf("The square has area: %f\n", area)
}
