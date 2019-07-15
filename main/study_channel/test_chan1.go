package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		c <- 0
		group.Done()
	}()
	group.Wait()
	fmt.Println(<-c)
	time.Sleep(1000)
}
