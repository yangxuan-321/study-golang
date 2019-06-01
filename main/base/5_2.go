package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"study-golang/main/base/queue"
	"time"
)

// -----------------------接口的实现和定义-------------------------

// -- go 语言 -- 的接口是由使用者来定义
// -- 传统语言 -- 的接口是由实现者来定义
// GO语言里面 的接口的 概念 就像一个大黄鸭 。 使用者 觉得 是大黄鸭 就是 大黄鸭 。 否则就不是
// 接口的实现是隐式的
// 接口不会用到 接口的指针
// 因为接口内部本身就维护了一个指针

// ---------------
//  接口 的 变量包含了 两个东西:
//	1.实现者的类型
//	2.实现者的指针 -- 实现者的指针指向了实现者

// 所以 我们 最好 接口的实现方法时候 最好不要用指针 。  因为 当 实现者的指针-> 指向实现者的时候 。 因为实现者是个指针 就会造成  指针-> 指向 一个指针 。 指针的指针

/**
总结：
	1. 接口变量自带指针
	2. 接口变量同样采用值传递，几乎不需要使用接口的指针
	3. 指针接受者实现只能以指针的方式使用，值接受者都可以
*/
// ---------------
/**
表示任何类型 interface{} Go语言的任何类型表示法
*/
// -----------------------接口的实现和定义-------------------------

type Retriever interface {
	Get(url string) string
}

type RetrieverImpl struct {
	Constents string
}

//值接受者
func (r RetrieverImpl) Get(url string) string {
	return r.Constents
}

type RetrieverImpl1 struct {
	UserAgent string
	TimeOut   time.Duration
}

//指针接受者
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
	var r Retriever

	//此处 两者皆可 写不写 指针都行。
	//因为: 指针接受者实现只能以指针的方式使用, 接受者都可以
	r = RetrieverImpl{Constents: "this is a fake imooc.com"}
	r = &RetrieverImpl{Constents: "this is a fake imooc.com"}

	inspect(r)
	//fmt.Println(download(r))

	// 此处只能 r = &RetrieverImpl1{} 如果 写成 r = RetrieverImpl1{} 会报错 。 取地址 因为 ： （指针接受者实现只能以指针的方式使用）
	r = &RetrieverImpl1{}
	//fmt.Println(download(retriever1))

	inspect(r)

	//强转 不是同一个类型 不能强转
	//所以要先判断是否能强转
	if rImpl, ok := r.(RetrieverImpl); ok {
		fmt.Println(rImpl.Constents)
	} else {
		fmt.Println("this object type is not RetrieverImpl, can not convert...")
	}

	//支持任何类型的 Queue 使用 interface{}
	queue := queue.QueueEveryObject{}
	fmt.Println("queue is empty:", queue.IsEmpty())

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push("yangxuan")
	fmt.Println("queue is empty:", queue.IsEmpty())

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case RetrieverImpl:
		fmt.Println("Contents:", v.Constents)
	case *RetrieverImpl1:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
