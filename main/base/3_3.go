package main

import "fmt"

func printSliceProp(slice []int) {
	fmt.Printf("len: %d, cap: %d \n", len(slice), cap(slice))
}

func main() {
	//定义一个切片  初始值 为 nil
	var s []int
	for i := 0; i < 100; i++ {
		printSliceProp(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	//切片的定义1 (先定义了一个arr装有0, 1, 2, 3 然后 又定义了一个 slice1 指向了 这个arr的视图)
	slice1 := []int{0, 1, 2, 3}
	printSliceProp(slice1)

	//只创建 size 为多大的 slice ， 并不初始化值。如何创建 且看
	//使用make内建函数 来创建一个 16大小容量的 int的数组 并分配试图给 slice2这个切片
	slice2 := make([]int, 16)
	printSliceProp(slice2)

	//创建一个 长度大小为 16 的切片 但是 指定 底层的数组大小 为 32
	slice3 := make([]int, 16, 32)
	printSliceProp(slice3)

	//-------------------------拷贝--------------------
	//将 slice1 的 内容 拷贝给 slice2
	copy(slice2, slice1)
	fmt.Println(slice2)

	//------------------------删除----------------------
	//例如删除 下标为 2 的元素
	//不能加法 例如 : slice2[:2] + slice2[:3]
	//所以想到 用 append 因为append函数的第二个参数 为 可变参数的 所以 需要加上 ...
	slice_23 := append(slice2[:2], slice2[3:]...)
	fmt.Println(slice_23)
	printSliceProp(slice_23)

	//删除 首尾
	fmt.Println(slice1)
	slice1 = slice1[1:]
	fmt.Println(slice1)
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1)
}
