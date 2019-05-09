package main

import (
	"fmt"
	"unicode/utf8"
)

//理解rune类型
//rune 相当于 int32  四个字节的

func testconvertrune(str string) {
	for _, v := range []rune(str) {
		fmt.Printf("%c   ", v)
	}
}

func main() {
	testconvertrune("我爱北京天安门")
	//如果想获得 字符的 个数
	//使用len是不行的 len只是获得字节的长度
	//需要使用 utf8.RuneCountInString()
	fmt.Println()
	fmt.Println(utf8.RuneCountInString("我爱北京天安门"))
	fmt.Println(len("我爱北京天安门"))

	//strings有很多string的工具类
	//strings.
}
