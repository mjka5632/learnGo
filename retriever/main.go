package main

import (
	"studygolang/retriever/mock"
	"studygolang/retriever/real"
	"fmt"
	"time"
)

//定义接口
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = &mock.Retriever{"This is a fake Baidu"}
	inspect(r)
	//fmt.Println(download(r))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	//fmt.Println(download(r))
	inspect(r)
	if sv, ok := r.(Retriever); ok {

		fmt.Printf("v implements String(): %s\n", sv.Get)
	}
	//realRetriever:=r.(*real.Retriever)
	//fmt.Println(realRetriever.TimeOut)
	//Type asseration
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(mockRetriever.Contents)
}

//判断类型
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
