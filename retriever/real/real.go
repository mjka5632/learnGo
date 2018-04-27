package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {

	resp, error := http.Get(url)

	if error != nil {
		panic(error)
	}
	//打印body信息，得到返回值
	result, error := httputil.DumpResponse(resp, true)
	//关闭连接
	resp.Body.Close()
	if error != nil {
		panic(error)
	}

	return string(result)
}
