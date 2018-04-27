package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

//转为二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		//数字转字符串
		result = strconv.Itoa(lsb) + result

	}
	return result
}

//扫描读每一行文件
func printFile(filenames string) {
	file, error := os.Open(filenames)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

//死循环
func forver() {
	for {

		fmt.Println("死循环")
	}

}
func main() {
	//forver()
	printFile("abc.txt")
	fmt.Println(
		convertToBin(15),
		convertToBin(0),

	)

}
