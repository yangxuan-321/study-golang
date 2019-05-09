package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	//这里和C语言不同的是：C语言的写法应该是 root.right->left 因为 root.right是个指针 指针访问 某个属性需要用 -> 而节点 只需要 . 既可以
	//不论是地址 还是 结构体（节点）本身，一律使用 点 . 来访问成员
	//Go语言没有构造函数
	root.right.left = new(treeNode)
	fmt.Println(root)
	nodes := []treeNode{
		{value: 3},
		{},
		root,
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
