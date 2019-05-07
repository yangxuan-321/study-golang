package main

import "fmt"

//切片就像 数组 的 视图    改变了 切片里面的内容 就改变了 数组的 内容
//[] 方括号 里面不写值 就是 代表 切片的意思
//所以 切片 比 数组 要好用很多
//因为 数组 直接是值传递 想改数组内容 必须用 数组指针（太麻烦）
//所以 切片 就不错。GO语言用 切片 比 数组 多
func slicetest(slice1 []int) {
	for i := range slice1 {
		slice1[i] = 111
	}
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr[0:3])
	fmt.Println(arr[2:])
	fmt.Println(arr[:5])
	fmt.Println(arr[:])
	slicetest(arr[:])
	fmt.Println(arr)
	//slice 可以 多次 slice
	//例如
	s1 := arr[0:5]
	s2 := s1[3:]
	fmt.Println(s2)
	s3 := append(s2, 101)
	s4 := append(s3, 102)
	s5 := append(s4, 103)
	s6 := append(s5, 104)
	s7 := append(s6, 105)
	s8 := append(s7, 106)
	fmt.Println(s3, s4, s5, s6, s7, s8)
	fmt.Println(arr)
	//因为s2 的 容量是 2 索引是 [3, 5) 但是 由于 数组的 大小是 9 所以
	//可以被操作的 索引 就是 [3, 9) 但是 只能进行 切片操作 例如 s3 = s2[3:8] 不能进行如: s2[8] 的取值操作
	//在对视图进行append操作时候 会 从 视图不包括的 下一个元素 对 原数组 进行覆盖 例如 对s2进行append 第一个被覆盖的就是 s2[5] 依次 s2[6]...
	//如果 append 的 内容大于了 整个数组的大小， 就会重新拷贝一个新的数组 用于 后续的追加。原数组的大小始终不变

}
