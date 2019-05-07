package main

import "fmt"

//GO语言的传参方式
//GO语言 只有 值传递 方式
//但是可以通过 指针的 传递 来达到 引用传递的效果

//值传递 不能 完成交换
func swap(a, b int) {
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
}

func swap1(a, b *int) {
	//tmp := *a
	//*a = *b
	//*b = tmp
	*a, *b = *b, *a
}

//指针
func main() {
	a, b := 1, 2
	swap(a, b)
	fmt.Println(a, b)
	swap1(&a, &b)
	fmt.Println(a, b)
}
