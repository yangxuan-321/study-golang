package main

import "fmt"

//---------------寻找公共最长子串-----------------------
//--自己解法--
func my(str string) int {
	max_cnt := 0
	curr_cnt := 0
	maps := make(map[int32]int)
	for _, v := range str {
		if _, ok := maps[v]; ok {
			curr_cnt = 0
			maps = make(map[int32]int)
		} else {
			curr_cnt++
			maps[v] = 1
		}
		if max_cnt < curr_cnt {
			max_cnt = curr_cnt
		}
	}
	return max_cnt
}

//别人解法
func other(str string) int {
	lastSameValueIndex := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, v := range []rune(str) {
		index, ok := lastSameValueIndex[v]
		if ok && index >= start {
			//遇到 被相似的了
			start = index + 1
		}

		if maxLength < (i - start + 1) {
			maxLength = (i - start + 1)
		}

		lastSameValueIndex[v] = i
	}
	return maxLength
}

func main() {
	//fmt.Println(my("abcabcbb"))
	//fmt.Println(my("bbbbb"))
	//fmt.Println(my("pwwkew"))

	fmt.Println(other("abcabcbb"))
	fmt.Println(other("bbbbb"))
	fmt.Println(other("pwwkew"))
	fmt.Println(other("一二三二一"))
}
