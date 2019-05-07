package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	//复数 1i == i  i*i = -1
	//演示 欧拉公式 ， 虚数部分不为0 是因为复数 由 64的实部 和 64位的虚部组成 除不尽
	fmt.Println(cmplx.Pow(math.E, math.Pi*1i) + 1)
}

//类型转化 -- GO语言的类型转换是强制的
func typeConvert() {
	var a int = 3
	var b int = 4
	var c int
	c = int(math.Sqrt(float64(a*a) + float64(b*b)))
	fmt.Println(c)
}

func main() {
	//euler()
	typeConvert()
}
