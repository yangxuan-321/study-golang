package main

import (
	"fmt"
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
			fmt.Printf("workid = %d, receive=%c", workId, <-w.in)
			fmt.Printf("\n")
			// 通知外面 打印结束
			w.done <- true
		}
	}()
	return w
}

type worker struct {
	in   chan int
	done chan bool
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
		workers[i].in <- 'A' + i //发数据
	}

	// 延时 5s 防止 主线程过快结束
	//time.Sleep(time.Second * 5)

	// 以上的代码延时 用于 主线程过快结束。 那么如何 能安全的 让其他线程 都做玩工作  然后 告知 主线程 再结束呢？
	// 1.使用Channel

}

func main() {
	chanDemo()
}
