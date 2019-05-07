package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertBin(num int) string {
	if 0 == num {
		return "0"
	}
	result := ""
	for ; num > 0; num /= 2 {
		result = strconv.Itoa(num%2) + result
	}
	return result
}

func printFile() {
	const filename = "./main/file/123.txt"
	file, err := os.Open(filename)
	if nil != err {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever() {
	//for ; ;  {
	//fmt.Println("abc")
	//}
	//等同于
	for {
		fmt.Println("abc")
	}

}

func main() {
	//fmt.Println(convertBin(0))
	printFile()
	//forever()
}
