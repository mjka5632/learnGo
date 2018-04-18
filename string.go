package main

import (
	"strings"
	"fmt"
)

func main() {
	origs := "Hi there!"
	news := ""
	//表示重复两次出现Hi there!
	news = strings.Repeat(origs, 2)
	fmt.Printf("The new repeated strign is:%s\n", news)

	//
	fmt.Println("-------split join------")
	str := "The quick brown fox jumps over the lazy dog"
	//转换为数组的形式
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice:%v\n", sl)
	//不想取返回参数，用_
	for _, val := range sl {
		fmt.Printf("%s -", val)
	}
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice:%v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s--", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
	//------常用转换方法
	//strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。
	//strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。
	//strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
	// 将 64 位浮点型的数字转换为字符串
	// 其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
}
