package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
)

// 使用 net/http/pprof 进行性能分析  为什么 前面 加上 "_" 是因为 如果不加 "_" 那么 因为 这个包 没有被现实的用，在ctrl+s 保存之后。就会自动的给删掉 导入的包。
// 然后 来 根据 分析的 情况 来优化程序。
// 可以通过 godoc -http :8888 命令来开启doc文档服务。 通过浏览器访问 http://localhost:8888
func main() {
	//resp, err := http.Get("http://www.imooc.com")
	var url string = "http://www.imooc.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("enter, some url be redirect:", *req)
			return nil
		},
	}
	//resp, err := http.DefaultClient.Do(request)
	resp, err := client.Do(request)
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if nil != err {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}
