package main

var a = "G"

func n() {
	print(a)
}

func m() {
	a = "0"
	print(a)
}
func main() {
	n()
	//m函数修改过后全局变量a的值就成了0
	m()
	n()
}
