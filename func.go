package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

/**
函数相关
 */
//转换
func swap(a, b int) (int, int) {

	return b, a

}

//转换 指针
func swap2(a, b *int) {

	*a, *b = *b, *a

}

//switch计算机功能
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("Not found op:%s", op)

	}

}
func div(a, b int) (q, r int) {
	//q=a/b
	//r=a%b
	return a / b, a % b

}

//嵌套函数
func apply(op func(int, int) int, a, b int) int {
	//拿到opName
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)\n", opName, a, b)

	return op(a, b)

}
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]

	}
	return s

}

func main() {
	a, b := 5, 6
	//&用来取内存地址
	swap2(&a, &b)
	fmt.Println(a, b)

	if result, error := eval(4, 3, "]"); error != nil {
		fmt.Println("errot", error)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 4)
	fmt.Println(
		sum(1, 2, 3, 4, 5),
		//eval(3,4,"+"),
		q, r,

	)

	//fmt.Println(div(13, 3))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4, ))
}
