package main

import (
	"reflect"
	"fmt"
)

type TagType struct {
	f1 bool   "f1 ..."
	f2 string "f2 sss"
	f3 int    "f3 iiiiii"
}

//获取tag的方法
func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixFiled := ttType.Field(ix)
	fmt.Printf("%v\n", ixFiled.Tag)
}

func main() {

	tt := TagType{true, "abcd", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
	//a,i:=for range tt{
	//	refTag(a,i)
	//}

}
