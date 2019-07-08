package main

import (
	"fmt"
	"regexp"
)

const text = "my mail is yangxuan321@gmail.com " +
	"mail2 is yangxuan_212@gmail.com1"

func main() {
	// 和 Compile 基本差不多， 只不过 MustCompile 表示必须编译成功
	// 不能写成 ^[a-zA-Z0-9]+@gmail\\.[a-zA-Z0-9]+$ 因为 这样的话， 必须 是
	// yangxuan321@gmail.com 的字符串 能找到 向 my mail is yangxuan321@gmail.com 就不行
	// 因为 ^代表字符串必须以什么开始 $必须以什么结束
	//re := regexp.MustCompile("[a-zA-Z0-9_]+@gmail\\.[a-zA-Z0-9]+")
	re := regexp.MustCompile(`([a-zA-Z0-9_]+)@(gmail\.[a-zA-Z0-9]+)`)
	//s := re.FindString(text)
	// 最后 -1 的参数就意味找所有
	allString := re.FindAllString(text, -1)
	fmt.Println(allString)

	// 如果 我们想 分别将 @前 和 @后的内容分开。我们需要给 正则 表达式
	// [a-zA-Z0-9]+@gmail\\.[a-zA-Z0-9]+ 加上 () 变成 ([a-zA-Z0-9]+)@(gmail\\.[a-zA-Z0-9])
	matchs := re.FindAllStringSubmatch(text, -1)
	for _, m := range matchs {
		if len(m) < 2 {
			continue
		}
		fmt.Println(m[1:])
	}
}
