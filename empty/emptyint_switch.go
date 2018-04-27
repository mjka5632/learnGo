package main

import "fmt"

type specialString string

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
		default:
			fmt.Println("Unknown Type!")
		}
	}
	testFuc(whatIsThis)
}

func main() {

	TypeSwitch()
}
