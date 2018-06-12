package main

import (
	"strings"
	"fmt"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	addBmp :=MakeAddSuffix(".bmp")
	addJpeg:=MakeAddSuffix(".jpeg")
	addBmp("file")
	addJpeg("file")
	fmt.Print(addJpeg("abc"))
}
