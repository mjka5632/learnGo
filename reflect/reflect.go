package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64=3.4
	//通过反射获得x的类型
	fmt.Println("Type:",reflect.TypeOf(x))
	//通过反射得到x的值
	v := reflect.ValueOf(x)
	fmt.Println("Value:",v)
	fmt.Println("Type:",v.Type())
	fmt.Println("Kind:",v.Kind())
	fmt.Println("Value:",v.Float())
	fmt.Println("Interface:",v.Interface())
	fmt.Printf("value is %5.2e\n",v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

}
