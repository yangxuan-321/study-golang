package main

import (
	"fmt"
	"time"
)

//channel 作为 返回值
func createIntChannel() chan int {
	return make(chan int)
}

//channel 作为 参数
func receiver(workId int, c chan int) {
	for {
		//有时候会出现 换行来不及打印出来 就 打印了 另一条 。  连在一起了 这是 典型的线程安全问题
		fmt.Printf("workid = %d, receive=%c", workId, <-c)
		fmt.Printf("\n")
	}
}

//上面的receiver方法不知道 什么时候 channel发完， 因此 一直使用死循环 进行接收。其实实际中 是有 办法知道 发送者是否发完的
//发送方做个 close
func receiver1(c chan int) {
	for {
		ch, ok := <-c
		if !ok {
			break
		}
		fmt.Println(ch)
	}

	// 或者可以写成 range
	// 发完退出 就 跳出循环了
	//for ch := range c {
	//	fmt.Println(ch)
	//}
}

// channel 的理论基础 是基于 CSP (Communication Sequential Process)
// 不要通过共享内存来通信，而要通过通信来共享内存 (GO语言创始人)

func channelClose() {
	c := make(chan int, 3)
	go receiver1(c)
	c <- '1'
	c <- '2'
	close(c)
	time.Sleep(time.Second)
}

//规定类型 -- 返回的channel只能 存数据 指定方向
// chan<- 存数据
// <-chan 取数据
func createWorker(workId int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			//有时候会出现 换行来不及打印出来 就 打印了 另一条 。  连在一起了 这是 典型的线程安全问题
			fmt.Printf("workid = %d, receive=%c", workId, <-c)
			fmt.Printf("\n")
		}
	}()
	return c
}

const channelSize = 10

//函数式一等公民 channel也是一等公民
func chanDemo() {
	var channels [10]chan<- int // c == nil
	//c := make(chan int)
	//专门开一个Goroutine来接受数据
	//go func() {
	//	//不停地收
	//	for {
	//		var receive int
	//		receive = <- c
	//		fmt.Println(receive)
	//	}
	//}()

	//写法一
	//for i := 0; i < channel_size; i++ {
	//	//channels[i] = make(chan int)
	//	channels[i] = createIntChannel()
	//	go receiver(i, channels[i])
	//}

	//写法二
	for i := 0; i < channelSize; i++ {
		channels[i] = createWorker(i)
	}

	//不停的发
	//for i := 0; i < 10; i ++ {
	//	c <- 1	//发数据
	//	c <- 2	//发数据
	//}

	for i := 0; i < channelSize; i++ {
		channels[i] <- '0' + i //发数据
	}

	for i := 0; i < channelSize; i++ {
		channels[i] <- 'A' + i //发数据
		//以下两句会报错，因为 我们 创造的channel只能用于 存数据， 不能取数据
		// r := <- channels[i]
		// fmt.Println(r)
	}

	// 延时 5s 防止 主线程过快结束
	time.Sleep(time.Second * 5)
}

func bufferedChannel() {
	// 带缓冲的Channel
	// 如果不带缓冲必须要 指定 接口这者。
	// 缓冲区大小是3， 如果没有人接受的话， 只能放三个 放第四个就会deadline掉
	// 使用缓冲区在性能上是有一定 优势的。 不会造成快速的 协程切换
	c := make(chan int, 3)
	go receiver(7, c)
	c <- '1'
	c <- '2'
	//延时一段时间 可以 使协程 进行切换
	time.Sleep(time.Second)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
