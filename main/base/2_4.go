package main

import (
	"fmt"
	"io/ioutil"
)

func iftest() {
	const filename = "./main/file/123.txt"
	//context, err := ioutil.ReadFile(filename)
	//if err != nil{
	//	fmt.Println("出错了,异常信息: ", err)
	//}else {
	//	fmt.Println("文件内容: ", context)
	//}

	//等效于
	// if 可以 有多个语句
	if context, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("出错了,异常信息: ", err)
	} else {
		fmt.Println("文件内容: ", context)
	}
}

//switch 默认 会 会在 case 带上break， 和其他语言不同。 如果 不想自动break 需要添加 fallthrough
func switchtest(a, b int, operate string) int {
	var result int = 0
	switch operate {
	case "+":
		result = a + b
		fallthrough
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		//程序报错
		panic("unsupported operate:" + operate)
	}
	return result
}

func switchtest1(score int) string {
	result := "未知"
	switch {
	case score > 100 || score < 0:
		panic("warning socre:" + string(score))
	case score < 60:
		result = "不及格"
	case score < 70:
		result = "及格"
	case score < 80:
		result = "中"
	case score < 90:
		result = "中"
	case score <= 100:
		result = "优秀"
	}
	return result
}

func main() {
	//iftest()
	//result := switchtest(10, 20, "+")
	//fmt.Println(result)
	fmt.Println(switchtest1(10))
}
