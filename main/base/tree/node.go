package tree

import (
	"fmt"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//为结构体定义方法
//其实也相当于参数传递  只不过是 参数 写在 方法 之前
//也可以 理解成  一个 值传递  举例：可以理解成 有一个 对象是 node_  =>  (node_)print  =>  node_.print()
func (node_ TreeNode) Print() {
	fmt.Println(node_.Value)
}

//既然 是 值传递 那么下面的方法 则 肯定 是 不能实现 set的
func (node_ TreeNode) SetValue_(value int) {
	node_.Value = value
}

//使用 地址传递 传递指针 方法 可以实现 set 方法
//nil 指针 也能调用方法!
//使用指针传递时 在方法里面使用 传递的指针的 对象时  就可理解成 Java里面的 this
func (node_ *TreeNode) SetValue(value int) {
	//可以将 node_ 理解成 Java的this python的self C语言使用结构体里面的指针函数时那个东西
	if nil == node_ {
		fmt.Println("node_ should be not nil, this operation is ignored!")
		return
	}
	//如果在C语言 里面 就需要写成 node_->value = value  或者 (*node_).value = value
	node_.Value = value
}

//自定义 工厂函数
//以下代码放在C语言 肯定是错误的
//因为 在 C语言里面 凡是 没有经过 malloc 的变量都是在 栈 上  ，  malloc的内存才是在堆上
//所以 如果 直接 声明一个 变量的话 ， 声明周期 在出方法 就已经 被销毁了。 导致 返回值 是一个 废弃的 内存
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

//遍历二叉树
func (node_ *TreeNode) Traverse() {
	if nil == node_ {
		return
	}
	node_.Left.Traverse()
	node_.Print()
	node_.Right.Traverse()
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
