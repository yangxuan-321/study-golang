package main

import (
	"fmt"
	"sync"
)

//规定类型 -- 返回的channel只能 存数据 指定方向
// chan<- 存数据
// <-chan 取数据
func createWorker(workId int) worker {
	w := worker{
		make(chan int),
		make(chan bool),
	}
	go func() {
		for {
			//有时候会出现 换行来不及打印出来 就 打印了 另一条 。  连在一起了 这是 典型的线程安全问题
			fmt.Printf("workid = %d, receive=%c \n", workId, <-w.in)
			// 通知外面 打印结束
			w.done <- true
		}
	}()
	return w
}

func createWorker1(workId int, wg *sync.WaitGroup) worker1 {
	w := worker1{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go func() {
		for {
			//有时候会出现 换行来不及打印出来 就 打印了 另一条 。  连在一起了 这是 典型的线程安全问题
			fmt.Printf("workid = %d, receive=%c \n", workId, <-w.in)
			// 通知外面 打印结束
			w.done()
		}
	}()
	return w
}

type worker struct {
	in   chan int
	done chan bool
}

type worker1 struct {
	in   chan int
	done func()
}

//函数式一等公民 channel也是一等公民
func chanDemo() {
	var workers [10]worker // c == nil
	//写法二
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- '0' + i //发数据
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i //发数据
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done
	}

	// 延时 5s 防止 主线程过快结束
	//time.Sleep(time.Second * 5)

	// 以上的代码延时 用于 主线程过快结束。 那么如何 能安全的 让其他线程 都做玩工作  然后 告知 主线程 再结束呢？
	// 1.使用Channel

}

//函数式一等公民 channel也是一等公民
func chanDemo1() {
	var workers [10]worker1 // c == nil
	wg := sync.WaitGroup{}
	wg.Add(10)
	//写法二
	for i := 0; i < 10; i++ {
		workers[i] = createWorker1(i, &wg)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- '0' + i //发数据
	}

	wg.Wait()
}

func ccc() {
	is := make(chan int)
	group := sync.WaitGroup{}
	group.Add(20)
	go func() {
		for {
			fmt.Println(<-is)
		}
	}()
	for i := 0; i < 20; i++ {
		is <- i
	}

}

// deadLock 因为 Channel 在发送和接受消息时 都是 会挂起  当前 线程都处于 挂起 状态。除非 另一端 转备好，否则 goroutine 都处于 挂起状态。
func ccc1() {
	is := make(chan int)
	//go func() {
	//	fmt.Println(<- is)
	//}()
	is <- 1
}

// 也会出现 deadlock
func ccc2() {
	is := make(chan int)
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		fmt.Println(<-is)
		group.Done()
	}()
	group.Wait()
}

func main() {
	//chanDemo1()
	ccc2()
}
