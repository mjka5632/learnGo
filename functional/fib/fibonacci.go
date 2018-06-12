package fib

func Fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		//调用本身用一直到<=1每次叠加
		res = Fibonacci1(n-1) + Fibonacci1(n-2)
	}
	return
}
