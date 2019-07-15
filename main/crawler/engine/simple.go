package engine

import (
	"log"
	"study-golang/main/crawler/fetcher"
)

type SimpleEngine struct {
}

// seeds 代表 初始的 种子， 也就是 最初始 的url

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			continue
		}

		// 加 三个 点， 就可以将 result.Requests 展开 加进去
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("fetch url:%s", r.Url)

	// 先获取 url 对应的 网址
	body, err := fetcher.Fetch(r.Url)
	if nil != err {
		// 如果发生了 错误 ， 打印错误， 继续下一次请求
		log.Printf("Fetcher url:%s error: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body)
}
