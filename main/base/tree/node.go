package main

import (
	"fmt"
)

type treeNode struct {
	value       int
	left, right *treeNode
}

//为结构体定义方法
//其实也相当于参数传递  只不过是 参数 写在 方法 之前
//也可以 理解成  一个 值传递  举例：可以理解成 有一个 对象是 node_  ->  (node_)print  ->  node_.print()
func (node_ treeNode) print() {
	fmt.Println(node_.value)
}

//既然 是 值传递 那么下面的方法 则 肯定 是 不能实现 set的
func (node_ treeNode) setValue_(value int) {
	node_.value = value
}

//使用 地址传递 传递指针 方法 可以实现 set 方法
func (node_ *treeNode) setValue(value int) {
	//如果在C语言 里面 就需要写成 node_->value = value  或者 (*node_).value = value
	node_.value = value
}

//自定义 工厂函数
//以下代码放在C语言 肯定是错误的
//因为 在 C语言里面 凡是 没有经过 malloc 的变量都是在 栈 上  ，  malloc的内存才是在堆上
//所以 如果 直接 声明一个 变量的话 ， 声明周期 在出方法 就已经 被销毁了。 导致 返回值 是一个 废弃的 内存
func createNode(value int) *treeNode {
	return &treeNode{value: value}
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
	root.left.right = createNode(3)
	fmt.Println(root)
	nodes := []treeNode{
		{value: 3},
		{},
		root,
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.print()

	//将 root的value 改为 100
	root.value = 100
	//打印 发现 修改成功
	root.print()

	//测试 set 方法
	root.setValue_(50)
	//发现并未修改成功 仍然是 100。 因为是值传递而已
	root.print()

	//测试 新的 set方法 使用指针传递
	//能修改成功

}
