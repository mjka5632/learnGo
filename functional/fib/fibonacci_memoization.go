package fib

const LIM = 41

var fibs [LIM]uint64
//缓存
func FibonacciMe(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = FibonacciMe(n-1) + FibonacciMe(n-2)
	}
	fibs[n] = res
	return
}
