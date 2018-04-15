package main

import "fmt"
var (
	cc="cc"
	ss=3
	status=true
)

//默认值
func variableZeroValue(){
	var s string//""
	var a int// 0
	//打印默认值 必须用Printf
	// %q是指字符串 %d是指数字
	fmt.Printf("%q %d",s,a)
}
//变量声明方式-【1】
func variableInitialValue(){
	var a,b int =3,4
	var s string="abc"
	fmt.Println(a,b,s)
}
//[2]
func variableTypeDeduction(){
	var a,b,c,d=1,"String",true,"cc"
	fmt.Println(a,b,c,d)
}
//[3]推荐用 :=的方式（不能用于定义全局变量）
func variableShorter()  {
	//:= 省略了var关键字
	a,b,c,d:="aa",true,33,"bb"

	b=false

	fmt.Println(a,b,c,d)
}
//常量和枚举
func enums(){
	//iota实现自增值默认从0开始
	const(
		cpp = iota
		java
		python
		javaScript
	)
	fmt.Println(cpp,javaScript,python,java)
}

func main() {
fmt.Println("Hello Wlord")
enums();
	//variableShorter();
	//variableZeroValue()
//variableTypeDeduction()
}
