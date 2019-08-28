package main

import (
	"study-golang/main/crawler/engine"
	"study-golang/main/crawler/persist"
	"study-golang/main/crawler/scheduler"
	"study-golang/main/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemServer(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
