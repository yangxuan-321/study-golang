package main

//
//import (
//	"bufio"
//	"fmt"
//	"golang.org/x/net/html/charset"
//	"golang.org/x/text/encoding"
//	"golang.org/x/text/transform"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"regexp"
//)
//
//func main() {
//	resp, err := http.Get("http://www.zhenai.com/zhenghun")
//	if nil != err {
//		panic(err)
//	}
//	//关闭输出流
//	defer resp.Body.Close()
//
//	if http.StatusOK != resp.StatusCode {
//		fmt.Println("http get is error")
//		return
//	}
//
//	// 给到1024个字节的内容，可以用此方法，去猜测当前网页编码。因为如果你想用 所谓的 <meta charset = xxx> 来正则得到，不一定准。
//	encoding := determineEncoding(resp.Body)
//
//	// 进行编码转换
//	utf8Reader := transform.NewReader(resp.Body, encoding.NewDecoder())
//
//	bytes, err := ioutil.ReadAll(utf8Reader)
//	/**
//	解决乱码问题:
//		需要安装两个库
//		先安装gopm工具
//		go get -v -u github.com/gpmgo/gopm
//		在安装关于编码的两个库
//		gopm get -g -v golang.org/x/text
//		gopm get -g -v golang.org/x/net/html
//	*/
//	if nil != err {
//		panic(err)
//	}
//
//	printCityList(bytes)
//
//	//fmt.Printf("%s", bytes)
//}
//
//func determineEncoding(r io.Reader) encoding.Encoding {
//	// 如果使用Peek方法直接去read 1024 个字节。就不好，因为你读过了，read的指针就指向了1024，
//	// 下一次正真使用内容时，就读取不到了。因此 需要对 Reader 进行包装
//	bytes, e := bufio.NewReader(r).Peek(1024)
//	if nil != e {
//		panic(e)
//	}
//
//	encoding, _, _ := charset.DetermineEncoding(bytes, "")
//	return encoding
//}
//
//func printCityList(contents []byte) {
//	// 使用 `` 可以 保证 不会涉及到复杂的 转义。 例如 如果 用 "" 那么 在
//	// 在正则里面 需要 用到 . 那么就要用 \\. 而在`` 只需要 \.
//	// ^> 代表 只要不是 右括号
//	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//	all := compile.FindAllSubmatch(contents, -1)
//	for _, a := range all {
//		fmt.Printf("url:%s city:%s\n", a[1], a[2])
//	}
//
//	// 找到 470 个
//	fmt.Printf("count:%d", len(all))
//}

func gg() {

}
