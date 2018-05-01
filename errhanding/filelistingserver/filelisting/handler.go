package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}
//显示文件内容
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	//URL地址是否是/list/开头的判断
	//结果为：1,-1,0只有当0时符合格式
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with" + prefix)
	}
	//取出地址栏list后面的值
	path := request.URL.Path[len(prefix):]
	//打开
	file, error := os.Open(path)
	if error != nil {
		//http.Error(writer,
		//	error.Error(),
		//	http.StatusInternalServerError)
		return error
		//panic(error)
	}
	//defer函数在结束时发生
	defer file.Close()
	//读取
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	//写入
	writer.Write(all)

	return nil
}
