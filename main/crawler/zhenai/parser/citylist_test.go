package parser

import (
	"fmt"
	"study-golang/main/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	bytes, e := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if nil != e {
		panic(e)
	}

	fmt.Println(string(bytes))
	//ParseCityList()
}
