package main

import (
	"io/ioutil"
	"study-golang/main/crawler/zhenai/parser"
	"testing"
)

func TestParseCityList(t *testing.T) {
	bytes, e := ioutil.ReadFile("./citylist_test_data.html")

	if nil != e {
		panic(e)
	}

	result := parser.ParseCityList(bytes)
	//fmt.Println(string(bytes))
	const citySize = 470
	// 网格驱动测试
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"City: 阿坝",
		"City: 阿克苏",
		"City: 阿拉善盟",
	}

	if len(result.Requests) != citySize {
		t.Errorf("result should be have %d requests, but act size:%d", citySize, len(result.Requests))
	}

	if len(result.Items) != citySize {
		t.Errorf("result should be have %d items, but act size:%d", citySize, len(result.Requests))
	}

	// 验证url
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted the %d url:%s, but was:%s", i, url, result.Requests[i].Url)
		}
	}

	for i, item := range expectedCities {
		if result.Items[i] != item {
			t.Errorf("excepted the %d item:%s, but was:%s", i, item, result.Items[i])
		}
	}
}
