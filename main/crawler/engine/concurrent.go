package engine

import (
	"fmt"
	"study-golang/main/crawler/filter"
)

type ConcurrentEngine struct {
	// 执行引擎
	Scheduler Scheduler
	// worker数量
	WorkerCount int
	// 用来存放需要存储的item
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var itemCount = 0

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		// 打印item
		for _, item := range result.Items {
			itemCount++
			//fmt.Printf("Got %d item: %v\n", itemCount, item)
			go func() {
				e.ItemChan <- item
			}()
		}
		// 将request
		for _, request := range result.Requests {
			// 去重操作
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// 使用 布隆过滤器 存在true
var fl filter.BloomFilter = filter.NewMemoryBloomFilter(64<<20, 5)

func isDuplicate(url string) bool {
	if fl.HasString(url) {
		// 存在
		return true
	}
	// 不存在
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//in := make(chan Request)
			// tell schheduler im ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if nil != err {
				continue
			}
			out <- result
		}
	}()
}
