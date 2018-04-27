package main

import (
	"io/ioutil"
	"fmt"
)

//根据分数进行评级
func grade(score int) string {
	//初始化评级值
	g := ""
	switch {
	case score > 100 || score < 0:
		panic(fmt.Sprintf("Wrong score:%d", score))

	case score < 60:
		g = "F"
	case score < 80:
		g = "B"
	case score > 90:
		g = "A"
	}
	return g
}
func main() {

	fmt.Println(grade(30))

	const filename = "abc.txt"
	if contends, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s/n", contends)

	}
	//if err !=nil{
	//	fmt.Println(err)
	//}else {
	//	//用Printf("%s/n",)否则将打印的时byte数组
	//	fmt.Printf("%s/n",contends)
	//}
	//
}
