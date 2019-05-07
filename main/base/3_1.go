package main

import "fmt"

func arr_test1() {
	var a [10]int
	//赋初值  如果 以 := 这种形式 声明变量 必须 给数组 赋初始值
	b := [3]int{1, 2, 3}
	//...代表让编译器帮我们计算个数
	c := [...]int{1, 2, 3, 4, 5, 6}
	//二维数组 3行4列
	var d [3][4]int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//数组的遍历 range 的使用
	for i, v := range b {
		// i 是下标  v 是值
		fmt.Println(i, v)
	}

	// _ 是为了 省略变量
	for _, v := range b {
		// i 是下标  v 是值
		fmt.Println(v)
	}
}

//数组是值类型
func printArray(arr [3]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
	arr[0] = 100
}

//数组指针
func pointArray(arr *[3]int) {
	for _, v := range arr {
		fmt.Println(v)
	}
	arr[0] = 100
}

//GO语言一般不适用 数组 一般使用 切片
//数组练习   [10]int 和 [20]int
func main() {
	//arr_test1()
	arr := [3]int{1, 2, 3}
	printArray(arr)
	printArray(arr)
	pointArray(&arr)
	pointArray(&arr)
}
