package main

import "fmt"

type File struct {
	fd   int
	name string
}
//工厂方法创建
func NewFile(fd int, name string) *File {

	if fd < 0 {
		return nil
	}
	fl := new(File)
	fl.fd = fd
	fl.name = name
	return fl //&File{fd,name}

}

func main() {
	f := NewFile(10, "./test.txt")
	//f:NewFile()
	fmt.Println(f)
}
