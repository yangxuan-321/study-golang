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

func createWorker(workId int) chan int {
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

const channel_size = 10

//函数式一等公民 channel也是一等公民
func chanDemo() {
	var channels [10]chan int // c == nil
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
	for i := 0; i < channel_size; i++ {
		channels[i] = createWorker(i)
	}

	//不停的发
	//for i := 0; i < 10; i ++ {
	//	c <- 1	//发数据
	//	c <- 2	//发数据
	//}

	for i := 0; i < channel_size; i++ {
		channels[i] <- '0' + i //发数据
	}

	for i := 0; i < channel_size; i++ {
		channels[i] <- 'A' + i //发数据
	}

	// 延时 5s 防止 主线程过快结束
	time.Sleep(time.Second * 5)
}

func main() {
	chanDemo()
}
