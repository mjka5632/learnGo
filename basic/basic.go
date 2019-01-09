package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

//变量命名规则遵循骆驼命名法，首个字母小写，其余单词首字母大写【外部包所用需要首字母也大写】
//全局变量 允许声明但不使用
var (
	cc     = "cc"
	ss     = 3
	status = true
	//[+]可以连接字符
	hi = "hel" + "lo"
)

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64((a*a + b*b))))
	return c
}

//常量和枚举
func enums() {
	//iota实现自增值默认从0开始
	const (
		cpp        = iota
		java
		python
		javaScript
	)
	fmt.Println(cpp, javaScript, python, java)
}

//默认值
func variableZeroValue() {
	var s string //""
	var a int    // 0
	//打印默认值 必须用Printf
	// %q是指字符串 %d是指数字
	fmt.Printf("%q %d", s, a)
}

//变量声明方式-【1】
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//[2]
func variableTypeDeduction() {
	var a, b, c, d = 1, "String", true, "cc"
	fmt.Println(a, b, c, d)
}

//[3]推荐用 :=的方式（不能用于定义全局变量）
func variableShorter() {
	//:= 省略了var关键字
	a, b, c, d := "aa", true, 33, "bb"

	b = false

	fmt.Println(a, b, c, d)
}

func main() {
	var goos string = runtime.GOARCH
	fmt.Printf("The operating system is: %s\n", goos)
	//通过os函数获取环境变量的路径值
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	//%v使用默认格式的标识符【结果相同】
	fmt.Printf("Path is %v\n", path)
	enums();
	//variableShorter();
	variableZeroValue()
	//variableTypeDeduction()
	fmt.Println(hi)
}
