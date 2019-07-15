package engine

import "fmt"

type ConcurrentEngine struct {
	// 执行引擎
	Scheduler Scheduler
	// worker数量
	WorkerCount int
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
			fmt.Printf("Got %d item: %v\n", itemCount, item)
		}
		// 将request
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
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
