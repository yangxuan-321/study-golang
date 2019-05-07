package main

import "fmt"

//在外面 使用 := 开声明变量 不行(报错)
//test3 := 1
// test2 不是一个 全局变量 因为只是 一个包类变量 作用域在包内部
var test2 = "ccc"
var test3 = "cc1"

//test3 和 test2 的变量定义 可以等价于
var (
	test4 = "cc2"
	test5 = "cc3"
)

func test1() {
	var name, age = "yangxuan", 18
	fmt.Println(name, age)
	var name1 string = "mina"
	var age1 int = 18
	fmt.Println(name1, age1)
	name2, age2 := "yang_mi", 18
	fmt.Println(name2, age2)
}

func main() {
	//fmt.Print("hello world!")
	test1()
	fmt.Println(test4, test5)
}
