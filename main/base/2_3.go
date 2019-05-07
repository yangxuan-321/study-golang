package main

import (
	"fmt"
	"math"
)

//常量学习
func constantsTest() {
	//这里的常量 可以 理解成 C语言 里面的 宏操作 只是相当于 一个 替换动作而已
	const filename = "123.txt"
	const a, b = 3, 4
	var c int
	// 因为 sqrt 需要的是 个 float64 类型的值 因此，这个地方不用强转 是因为是常量
	//const 数值可以作为各种类型使用
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp        = 0
		java       = 1
		python     = 2
		javascript = 3
	)
	//等同于
	const (
		chinese = iota
		math
		english
	)
	// iota 也可以 作为一个自增值 的 种子 只要能写出表达式 就可以
	const (
		b  = 1 << (10 * iota) //1 * 2的10*0次方		==			1 向后移动10位 10 000 000 000
		kb                    //1 * 2的10*1次方
		mb                    //1 * 2的10*2次方
		gb                    //...
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
