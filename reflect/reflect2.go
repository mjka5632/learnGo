package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64=3.4
	v:=reflect.ValueOf(x)
	//setting a value
	//v.SetFloat(3.111)
	//判断是否可以设置值
	fmt.Println("settability of v:",v.CanSet())
	v=reflect.ValueOf(&x)
	fmt.Println("type of v:",v.Type())
	fmt.Println("settability of v",v.CanSet())
	v=v.Elem()
	fmt.Println("The Elem of v is :",v)
	fmt.Println("settability of v:",v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
