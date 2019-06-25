package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(workId int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			//有时候会出现 换行来不及打印出来 就 打印了 另一条 。  连在一起了 这是 典型的线程安全问题
			fmt.Printf("workid = %d, receive=%d", workId, <-c)
			fmt.Printf("\n")
		}
	}()
	return c
}

func main() {
	//var c1, c2 chan int
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	var values []int

	n := 0
	//hasValue := false
	// 不用 select 以前对 channel 收发数据都是 会被阻塞的。
	// 使用了channel 就像 是 非阻塞的
	// 如果 不加 default 程序 还是 会 deadlock
	// 给 程序 计时 10秒
	// 使用 After 函数 会返回给 一个 channel。 到了 10 秒钟之后， 会自动往 channel 送一个 值
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			//fmt.Println("this value from c1. ", n)
			//hasValue = true
			values = append(values, n)
		case n = <-c2:
			//fmt.Println("this value from c2. ", n)
			//hasValue = true
			values = append(values, n)
		case activeWorker <- activeValue:
			//hasValue = false
			values = values[1:]
		// 超时（两次数据获取的时间 大与 800ms 就会超时）
		// 因为 select 中 新来的 case 会打断 ??? 不确定
		// 因为 走到 这个 case 的时候，会被阻塞 直到 新来的 select来打断他。或者 到了 0.5秒 就会自动 接收到 chan 的 一个值。 就完成了所谓的 超时判断问题
		case <-time.After(500 * time.Millisecond):
			fmt.Println("time out...")
		// 定时 打印 切片长度， 防止数据 积压 过多
		case <-tick:
			fmt.Println("slice length: ", len(values))

		case <-tm:
			fmt.Println("bye bye...")
			return
			//default:
			//	fmt.Println("no value received.")
		}
	}
	//time.Sleep(10000 * time.Millisecond)
}
