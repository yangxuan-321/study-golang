package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

// -----------------------接口的实现和定义-------------------------

// -- go 语言 -- 的接口是由使用者来定义
// -- 传统语言 -- 的接口是由实现者来定义
// GO语言里面 的接口的 概念 就像一个大黄鸭 。 使用者 觉得 是大黄鸭 就是 大黄鸭 。 否则就不是

// -----------------------接口的实现和定义-------------------------

type Retriever interface {
	Get(url string) string
}

type RetrieverImpl struct {
	Constents string
}

func (r *RetrieverImpl) Get(url string) string {
	return r.Constents
}

type RetrieverImpl1 struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *RetrieverImpl1) Get(url string) string {
	if resp, err := http.Get(url); nil == err {
		result, err := httputil.DumpResponse(resp, true)
		//用完关掉
		resp.Body.Close()
		if nil == err {
			return string(result)
		}
		panic("retriever error1")
	}
	panic("retriever error2")
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	fmt.Println(download(&RetrieverImpl{Constents: "this is a fake imooc.com"}))
	var retriever1 *RetrieverImpl1 = &RetrieverImpl1{}
	fmt.Println(download(retriever1))
}
