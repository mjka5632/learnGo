package main

func main() {

	var arr = [5]int{0, 1, 2, 3, 4}
	sum1(arr[:])
}

func sum1(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
