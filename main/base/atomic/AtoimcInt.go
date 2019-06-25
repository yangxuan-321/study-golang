package main

import (
	"fmt"
	"sync"
)

/**
 * 用 mutex
 * 做一个原子化的 int
 * go run -race ./main/base/atomic/AtoimcInt_.go 可以用来检测发现 程序中 数据访问冲突。
 */

type AtoimcInt_ int

type AtoimcInt struct {
	value int
	// 互斥锁
	lock sync.Mutex
}

func (a *AtoimcInt_) increment_() {
	*a++
}

func (a *AtoimcInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *AtoimcInt_) get_() int {

	return int(*a)
}

func (a *AtoimcInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	// 系统安全的 int32 相加操作
	//atomic.AddInt32()
	//var a *AtoimcInt_ = new(AtoimcInt_)
	var a *AtoimcInt = new(AtoimcInt)
	a.value = 0
	group := sync.WaitGroup{}
	group.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			//a.increment_()
			//fmt.Println(*a)

			a.increment()
			fmt.Println(a.get())

			group.Done()
		}()
	}
	group.Wait()
	// 如果 使用 传统的 加法 会发现 a 有可能 不等于 10000
	fmt.Println("-------------结束了: a = ", *a, "------------")
}
