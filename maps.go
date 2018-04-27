package main

import "fmt"

func main() {
	m := map[string]string{
		"Name":   "kim",
		"Course": "golang",
		"Boby":   "study",
		"ff":     "qq",
	}
	fmt.Println(m)
	fmt.Println("固定map的key的顺序")
	//定义一个切片g，存放map中的key
	g := []string{"Name", "Course", "Boby", "ff"}
	fmt.Println("slice取出来的", g)
	//通过变量g达到找到map中的value
	for _, v := range g {
		fmt.Println(m[v])
		//fmt.Println(v)
	}

	//fmt.Println(s)
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)
	//Map是无序的
	fmt.Println("Traversing map")
	for k := range m {
		fmt.Println(k)
	}
	fmt.Println("Getting values")
	//key不存在时打印空
	//ok代表判断是否存在key 存在：true 否则 ：false
	courseName, ok := m["Course"]
	fmt.Println(courseName, ok)
	//判断map中是否包含key
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("not exist key")
	}
	fmt.Println("deleting valus")
	r, ok := m["Name"]
	fmt.Println(r, ok)
	delete(m, "Name")
	a, ok := m["Name"]
	fmt.Println(a, ok)
}
