package main

import "fmt"

type specialString string
type flag bool

var tf flag
var whatIsThis specialString = "hello"


func TypeSwitch() {
	testFuc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("Type is %T\n", v)
		case string:
			fmt.Printf("Type is %T\n", v)
		case specialString:
			fmt.Printf("Type is %T\n", v)
		case flag:
			fmt.Printf("Type is %T\n", v)
		default:
			fmt.Println("Unknown Type!")
		}
	}
	testFuc(tf)
}

func main() {

	TypeSwitch()
}
