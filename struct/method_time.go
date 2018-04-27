package main

import (
	"time"
	"fmt"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	//String方法是把时间转换为String类型
	return t.Time.String()[0:3]
}

func main() {

	m := myTime{time.Now()}
	//调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	//调用first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}
