package main

import (
	"fmt"
	"study-golang/main/base/queue"
	"study-golang/main/base/tree"
)

//----------------------如何扩展已有类型---------------------------
// 1.如何扩充系统类型或者别人的类型
//		因为GO语言不支持继承所以只有两种方法
//		1).使用别名
//		2).使用组合
// 使用 组合 来 扩展 二叉树的 后后序遍历
// 使用 别名的方式 来 扩展数组 成为 队列 -- 详见 queue下面的 Queue.go的相关代码

type myTreeNode struct {
	node *tree.TreeNode
}

// 使用组合的 方式 来扩展 扩展后序遍历
func (myNode *myTreeNode) postOrder() {
	if nil == myNode || nil == myNode.node {
		return
	}
	var myNodeLeft *myTreeNode = &myTreeNode{node: myNode.node.Left}
	myNodeLeft.postOrder()
	var myNodeRight *myTreeNode = &myTreeNode{node: myNode.node.Right}
	myNodeRight.postOrder()
	myNode.node.Print()
}

func main() {
	root := tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{Value: 2}
	root.Right = &tree.TreeNode{Value: 4}
	var myNode *myTreeNode = &myTreeNode{node: &root}
	myNode.postOrder()

	var p *queue.Queue = new(queue.Queue)
	fmt.Println(p.IsEmpty())
	p.Push(1)
	p.Push(2)

	p.Push(3)
	fmt.Println(*p)
	p.Pop()
	fmt.Print(*p)
}
