package main

import (
	"net/http"
	"os"
	"log"
	"learnGo/errhanding/filelistingserver/filelisting"
)
//用户可见错误
type userError interface {
	error
	Message() string
}
//定义这个类型是为了 把http.HandleFunc的参数handler包装起来,
//返回一个error
type appHandler func(write http.ResponseWriter,
	request *http.Request) error

//处理错误的方法
func errWrapper(handler appHandler) func(
	write http.ResponseWriter,
	request *http.Request) {
		//首先返回函数本身
	return func(write http.ResponseWriter,
		request *http.Request) {
			//defer函数在结束时发生
		defer func() {
			//recover 仅仅在defer中调用不会因为错误，终断服务
			if r := recover(); r != nil {
				log.Printf("Panic %v", r)
				http.Error(write,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(write, request)

		if err != nil {
			//log.Warn("Error handing request %s",err.Error())
			log.Printf("Error handing request %s", err.Error())
			//判断是否属于用户可见错误
			if userErr, ok := err.(userError); ok {
				http.Error(write,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(write, http.StatusText(code), code)
		}
	}
}
func main() {
	//http.HandleFunc包含了(请求地址，handler)
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
	//a:="/list/list/list/aa"
	//index := strings.Index( a,"/list/")
	//fmt.Println(index)
}
