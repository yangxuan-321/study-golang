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
//也可以 理解成  一个 值传递  举例：可以理解成 有一个 对象是 node_  =>  (node_)print  =>  node_.print()
func (node_ treeNode) print() {
	fmt.Println(node_.value)
}

//既然 是 值传递 那么下面的方法 则 肯定 是 不能实现 set的
func (node_ treeNode) setValue_(value int) {
	node_.value = value
}

//使用 地址传递 传递指针 方法 可以实现 set 方法
//nil 指针 也能调用方法!
//使用指针传递时 在方法里面使用 传递的指针的 对象时  就可理解成 Java里面的 this
func (node_ *treeNode) setValue(value int) {
	//可以将 node_ 理解成 Java的this python的self C语言使用结构体里面的指针函数时那个东西
	if nil == node_ {
		fmt.Println("node_ should be not nil, this operation is ignored!")
		return
	}
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

//遍历二叉树
func (node_ *treeNode) traverse() {
	if nil == node_ {
		return
	}
	node_.left.traverse()
	node_.print()
	node_.right.traverse()
}

/**
 *	  ----------------------------------------------总结-------------------------------------------------
 *    func (node_ *treeNode) print(){}
 *		第一个括号里的 node_ *treeNode 被称为 函数接受者。接受者的类型和变量名都是我们 显式 定义的 。 函数接受者可以为指针
 *   也可以为值，但是切记 只有 指针 传递 才能到达 修改效果 改变结构内容
 *   				===================nil 指针 也能调用方法!================
 *            			指针接受者		vs			值接受者
 *					1.要改变结构体内容，必须使用指针接受者
 *					2.结构体过大，也考虑使用指针接受者
 *					3.一致性：如有指针接受者，最好都是指针接受者
 *
 *	  其实值接受者 才是 GO语言的 一个特性，因为指针接受者 向C++和Java都是类似的。例如 一个对象调用一个方法 。 那么在 对象方法
 * 里面 使用 this的 时候 就一定是这个对象的本身。不可能像值传递一样 帮你复制 一份。
 *   值/指针 接受者均可作接受 值/指针    可以理解成 来者不拒
 *   ----------------------------------------------总结-------------------------------------------------
 */

// --------------始终记住 值传递 就是 拷贝一份 数据 到新的对象 -------------------
// 值传递，改变的新对象的值，不会对原对象产生任何影响
// 地址传递，地址即指针，不会产生新的对象，会把原对象地址传过去，改变了什么，就是改变了原对象的什么
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
	root.setValue(70)
	//打印 70
	root.print()
	//如果说 ， 用的是指针传递， 不应该是 需要写成  (&root).setValue(70) 这样子嘛 吗？
	//其实 在 传递 调用者（参数）的 时候 会把 取完地址 自动 给到
	//但是以下写法 也没问题
	(&root).setValue(80)
	//打印 80
	root.print()

	//其实这样调用 print 方法也不会错
	//因为 如果 结构函数 要的是一个值 我们传的是 地址。他就会把地址 里面 值 取出来 交给print 函数
	//相对 func (node_ *treeNode) setValue(value int) 的方法 需要的是一个 指针(地址) ， 如果 我们传递给setValue一个值， 也
	//会自动帮我们转成 地址  。  以上 所说的 应该是 编译器 帮我们做的
	(&root).print()

	//测试 nil指针 也可以 调用方法
	//不给 nilNode 赋值  那么 nilNode 就会默认为 nil
	var nilNode *treeNode
	nilNode.setValue(110)
	//这个就会报错 因为函数接受者 不是一个指针。 只有nil指针才能调用方法
	//所以提醒我们 经常应该 使用 指针作为 函数接受者
	// nilNode.print()

	fmt.Println("================二叉树遍历=================")
	root.traverse()
	fmt.Println("================二叉树遍历=================")
}
