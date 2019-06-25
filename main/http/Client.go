package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//resp, err := http.Get("http://www.imooc.com")
	var url string = "http://www.imooc.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	defer resp.Body.Close()

	bytes, err := httputil.DumpResponse(resp, true)
	if nil != err {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}
