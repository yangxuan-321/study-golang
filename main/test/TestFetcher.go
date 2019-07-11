package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "http://album.zhenai.com/u/1373636812", nil)
	if nil != err {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	client := http.DefaultClient
	response, err := client.Do(request)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s", bytes)
	fmt.Println(err)
}
