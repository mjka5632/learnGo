package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func convertToBin(n int)  string{
	result:=""
	for ; n>0;n/=2  {
		lsb:=n%2
		//数字转字符串
		result=strconv.Itoa(lsb)+result

	}
	return result
}
func printFile(filenames string)  {
	file,error:=os.Open(filenames)
	if error!=nil{
		panic(error)
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}
func forver()  {
	for {

		fmt.Println("死循环")
	}

}
func main() {
	forver()
	printFile("abc.txt")
	fmt.Println(
		convertToBin(15),
		convertToBin(0),

	)


}
