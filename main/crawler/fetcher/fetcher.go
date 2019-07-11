package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"study-golang/main/crawler/util/http"
)

// 主要负责 网络请求
// 返回请求的内容(byte数组) 或者 相关错误
func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get(url)
	resp, err := httputils.Get(url)
	if nil != err {
		return nil, err
	}
	//关闭输出流
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		fmt.Println("http get is error")
		//return nil, errors.New("http get is error, statusCode:" + string(resp.StatusCode))
		return nil, fmt.Errorf("http get is error, statusCode: %d", resp.StatusCode)
	}

	// 给到1024个字节的内容，可以用此方法，去猜测当前网页编码。因为如果你想用 所谓的 <meta charset = xxx> 来正则得到，不一定准。
	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)

	// 进行编码转换
	utf8Reader := transform.NewReader(bodyReader, encoding.NewDecoder())

	bytes, err := ioutil.ReadAll(utf8Reader)
	/**
	解决乱码问题:
		需要安装两个库
		先安装gopm工具
		go get -v -u github.com/gpmgo/gopm
		在安装关于编码的两个库
		gopm get -g -v golang.org/x/text
		gopm get -g -v golang.org/x/net/html
	*/

	return bytes, err
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 如果直接去read 1024 个字节。就不好，因为你读过了，read的指针就指向了1024，
	// 下一次正真使用内容时，就读取不到了。因此使用Peek方法
	bytes, e := r.Peek(1024)
	if nil != e {
		//panic(e)
		log.Printf("Fetcher error, %v", e)
		//如果读取不成功，就返回默认的 utf-8 让接下来的程序 尝试读取
		return unicode.UTF8
	}

	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
