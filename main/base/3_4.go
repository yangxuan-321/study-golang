package main

import "fmt"

//Map
//1.map定义  xxxMap := map[key_type]value_type{}

func mapDefind1() {
	m := map[string]int{
		"yangxuan": 18,
		"mina":     17,
		"xxxxx":    11,
	}
	fmt.Println(m)

	//使用range遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	//通过key获得value
	fmt.Println(m["yangxuan"])
	//因为 GO语言 普通都有一个 ZERO VALUE的东西。所以 在key 不存在的时候 此例子的map会返回一个 空字符串
	//那我们如何判断 这个 对应的 这个 key是不存在 还是 key存在 value本来就是 空值呢？
	//请看
	if v, ok := m["xxxx"]; ok == true {
		fmt.Println(v)
	} else {
		fmt.Println("不存在的key")
	}

	//---------删除 map 元素 ------------
	_, ok := m["xxxxx"]
	fmt.Println(ok)
	delete(m, "xxxxx")
	_, ok = m["xxxxx"]
	fmt.Println(ok)
}

func mapDefind2() {
	//使用内建函数make创建的空map
	m := make(map[string]int) //实际值是emptymap
	fmt.Println(m)
	var m1 map[string]int //实际值是 nil  和 emptymap差不多
	fmt.Println(m1)
}

func main() {
	mapDefind1()
	mapDefind2()
}
