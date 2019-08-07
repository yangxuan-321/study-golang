package main

import (
	"fmt"
	"study-golang/main/crawler/filter"
)

var fl filter.BloomFilter = filter.NewMemoryBloomFilter(64<<20, 5)

func isDuplicate(url string) bool {
	if fl.HasString(url) {
		// 存在
		return true
	}
	// 不存在
	return false
}

func main() {
	fl.PutString("123")
	fl.PutString("456")
	fmt.Println(isDuplicate("123"))
}
