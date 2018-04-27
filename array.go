package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}
func main() {
	var arrKey = [5]string{3: "Chris", 4: "Ron"}
	for i, v := range arrKey {
		fmt.Println("arrkey %d %s/n", i, v)
	}
	//[；]每个都打印出来
	ar := [3]int{2, 3, 4}
	fmt.Println(ar[:])
	//[...]没有定义长度
	long := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(long)
	fmt.Println(ar[:])
	ar[1] = 5
	f(ar)
	fp(&ar)
	//定义长度为5的数组
	var arr1 [5]int
	/*for i:=0;i<len(arr1);i++{
		arr1[i]=i*2
	}*/
	a := [...]string{"a", "b", "c"}
	for i, v := range a {
		fmt.Println(i, v)
	}
	for i, _ := range arr1 {
		fmt.Println(i)

	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("index %d value %d\n", i, arr1[i])
	}

}
