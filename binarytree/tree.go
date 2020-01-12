package binarytree

import (
    "fmt"
    "github.com/hq-cml/data-structure/stack"
)

type TreeNode struct {
    Data interface{}
    Left *TreeNode
    Right *TreeNode
}

func NewTree(data interface{}, left, right *TreeNode) *TreeNode {
    return &TreeNode{
        Data:  data,
        Left:  left,
        Right: right,
    }
}

//递归，前序遍历
func(tree *TreeNode) PreOrder() {
    if tree == nil {
        return
    }
    fmt.Print(tree.Data, " ")
    tree.Left.PreOrder()
    tree.Right.PreOrder()
}

//递归，中遍历
func(tree *TreeNode) InOrder() {
    if tree == nil {
        return
    }
    tree.Left.InOrder()
    fmt.Print(tree.Data, " ")
    tree.Right.InOrder()
}

//递归，后遍历
func(tree *TreeNode) PostOrder() {
    if tree == nil {
        return
    }
    tree.Left.PostOrder()
    tree.Right.PostOrder()
    fmt.Print(tree.Data, " ")
}

//节点个数
func(tree *TreeNode) NodeCnt() int {
    if tree == nil {
        return 0
    }

    return 1 + tree.Left.NodeCnt() + tree.Right.NodeCnt()
}

//树高度
func(tree *TreeNode) TreeHeight() int {
    if tree == nil {
        return 0
    }

    return 1 + func(tree1, tree2 *TreeNode) int {
        if tree1.TreeHeight() > tree2.TreeHeight() {
            return tree1.TreeHeight()
        } else {
            return tree2.TreeHeight()
        }
    }(tree.Left, tree.Right)
}

//非递归，前序遍历
func(tree *TreeNode) PreOrderNoRecursion() {
    var myStack stack.Stack
    p := tree
    for p != nil || !myStack.IsEmpty() {
        for p != nil {
            fmt.Print(p.Data, " ")
            myStack.Push(p)
            p = p.Left
        }
        tmpNode, ok := myStack.Pop()
        if ok {
            tmp, _ := tmpNode.(*TreeNode)
            if tmp.Right != nil {
                p = tmp.Right
            }
        }
    }
}

//非递归，中序遍历
func(tree *TreeNode) InOrderNoRecursion() {
    var myStack stack.Stack
    p := tree
    for p != nil || !myStack.IsEmpty() {
        for p != nil {
            myStack.Push(p)
            p = p.Left
        }
        tmpNode, ok := myStack.Pop()
        if ok {
            tmp, _ := tmpNode.(*TreeNode)
            fmt.Print(tmp.Data, " ")
            if tmp.Right != nil {
                p = tmp.Right
            }
        }
    }
}

//非递归，后序遍历
//主要是利用了一个栈顶元素获取但不出栈的特性
func(tree *TreeNode) PostOrderNoRecursion() {
    var myStack stack.Stack
    p := tree
    var preNode *TreeNode
    for p != nil || !myStack.IsEmpty() {
        for p != nil {
            myStack.Push(p)
            p = p.Left
        }
        tmpNode, ok := myStack.Top() //只是获取栈顶元素，不出栈

        if ok {
            tmp, _ := tmpNode.(*TreeNode)
            if tmp.Right == nil || tmp.Right == preNode { //如果没有右孩子或右孩子已经访问过了，出栈
                fmt.Print(tmp.Data, " ")
                preNode = tmp
                myStack.Pop() //扔出元素
            } else {
                p = tmp.Right
            }
        }
    }
}

//兄弟孩子表示法
/*
 *        1
 *     2     3
 *   4   5  6  7
 *
 * 前序： 1 2 4 5 3 6 7
 * 中：   4 2 5 1 6 3 7
 * 后：   4 5 2 6 7 3 1
 *
 * 处理后：
 *         1
 *        2
 *     4    3
 *      5  6  7
 *
 * 前序： 1 2 4 5 3 6 7
 * 中：   4 5 2 6 7 3 1
 * 后：   5 4 7 6 3 2 1
 */
//非递归利用栈
func(tree *TreeNode)Change2ChildNoRecursion() {
    var myStack stack.Stack
    p := tree

    for p != nil || !myStack.IsEmpty() {
        for p != nil {
            myStack.Push(p)
            p = p.Left
        }
        tmpNode, _ := myStack.Top() //只是获取栈顶元素，不出栈
        tmp, _ := tmpNode.(*TreeNode)
        if tmp.Right == nil {
            myStack.Pop() //扔出元素
        } else {
            p = tmp.Right
            if tmp.Left == nil {
                tmp.Left = p
            } else {
                tmp.Left.Right = p
            }
            tmp.Right = nil   //Not forget to set right to NULL
        }
    }
}

//递归
func(tree *TreeNode)Change2Child() {
    if tree == nil {
        return
    }

    if tree.Right != nil {
        if tree.Left == nil {
            tree.Left = tree.Right
        } else {
            tree.Left.Change2Child()
            tree.Left.Right = tree.Right
        }
        tree.Right.Change2Child()
        tree.Right = nil
    } else {
        tree.Left.Change2Child()
    }
}







