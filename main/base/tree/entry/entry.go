package main

import (
	"fmt"
	"study-golang/main/base/tree"
)

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	//这里和C语言不同的是：C语言的写法应该是 root.right->left 因为 root.right是个指针 指针访问 某个属性需要用 -> 而节点 只需要 . 既可以
	//不论是地址 还是 结构体（节点）本身，一律使用 点 . 来访问成员
	//Go语言没有构造函数
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(3)
	fmt.Println(root)
	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		root,
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()

	//将 root的value 改为 100
	root.Value = 100
	//打印 发现 修改成功
	root.Print()

	//测试 set 方法
	root.SetValue_(50)
	//发现并未修改成功 仍然是 100。 因为是值传递而已
	root.Print()

	//测试 新的 set方法 使用指针传递
	//能修改成功
	root.SetValue(70)
	//打印 70
	root.Print()
	//如果说 ， 用的是指针传递， 不应该是 需要写成  (&root).setValue(70) 这样子嘛 吗？
	//其实 在 传递 调用者（参数）的 时候 会把 取完地址 自动 给到
	//但是以下写法 也没问题
	(&root).SetValue(80)
	//打印 80
	root.Print()

	//其实这样调用 print 方法也不会错
	//因为 如果 结构函数 要的是一个值 我们传的是 地址。他就会把地址 里面 值 取出来 交给print 函数
	//相对 func (node_ *treeNode) setValue(value int) 的方法 需要的是一个 指针(地址) ， 如果 我们传递给setValue一个值， 也
	//会自动帮我们转成 地址  。  以上 所说的 应该是 编译器 帮我们做的
	(&root).Print()

	//测试 nil指针 也可以 调用方法
	//不给 nilNode 赋值  那么 nilNode 就会默认为 nil
	var nilNode *tree.TreeNode
	nilNode.SetValue(110)
	//这个就会报错 因为函数接受者 不是一个指针。 只有nil指针才能调用方法
	//所以提醒我们 经常应该 使用 指针作为 函数接受者
	// nilNode.print()

	fmt.Println("================二叉树遍历=================")
	root.Traverse()
	fmt.Println("================二叉树遍历=================")
}
