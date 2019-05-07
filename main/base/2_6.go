package main

import (
	"fmt"
	"sort"
)

//多参数 单返回值
func evel(a, b int, operate string) int {
	result := 0
	switch operate {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsopported operate:" + operate)
	}
	return result
}

//多返回值
//func max_min(nums []int) (int, int) {
//等效于
func max_min(nums []int) (min, max int) {
	sort.Ints(nums)
	return nums[0], nums[len(nums)-1]
}

//等效于
func max_min1(nums []int) (min, max int) {
	sort.Ints(nums)
	min = nums[0]
	max = nums[len(nums)-1]
	return
}

//如何将多返回值的函数 -- 使用的时候 只关注某一个参数的场景
func max1(nums []int) int {
	//q, r := max_min(nums)
	//return r
	//以上代码 会报错 因为 q, r 是两个值 必须 都用到 才可以（因为GO语言定义到的变量都必须要用到）
	//采取以下 措施 就可以
	// 用 _ 代表参数 就代表 可以不被使用
	_, r := max_min(nums)
	return r
}

//多返回值 常被用于 异常的返回 毕竟 panic 直接 抛出异常 中断程序太暴力
//多参数 单返回值
func evel1(a, b int, operate string) (int, error) {
	result := 0
	var err error = nil
	switch operate {
	case "+":
		//result = a + b
		//函数指针
		result = opExe(add, a, b)
	case "-":
		result = a - b
	case "*":
		//匿名函数
		result = opExe(
			func(a1 int, b1 int) int {
				return a1 * b1
			},
			a, b)
	case "/":
		result = a / b
	default:
		//panic("unsopported operate:" + operate)
		err = fmt.Errorf("unsopported operate:" + operate)
	}
	return result, err
}

func add(a, b int) int {
	return a + b
}

//函数是编程 -- 函数本身也可以被当做一个参数 传入
func opExe(op func(int, int) int, a int, b int) int {
	return op(a, b)
}

// GO 语言 不能重载 重写 支持 可变参数列表
func unknownArgs(a ...int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i]
	}
	return result
}

func main() {
	//fmt.Println(evel(10, 20, "+"))
	//var nums = []int{8, 2, 4, 1, 7}
	//min, max := max_min(nums)
	//min, max := max_min1(nums)
	//fmt.Println(min, max)
	//fmt.Println(max1(nums))
	if result, err := evel1(10, 20, "+"); nil != err {
		fmt.Println("出现了异常:", err)
	} else {
		fmt.Println("结果为:", result)
	}

	if result, err := evel1(10, 20, "*"); nil != err {
		fmt.Println("出现了异常:", err)
	} else {
		fmt.Println("结果为:", result)
	}

	fmt.Println(unknownArgs(1, 2, 3))
}
