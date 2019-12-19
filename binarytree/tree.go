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
/*
 *        1
 *     2     3
 *   4   5  6  7
 */
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










