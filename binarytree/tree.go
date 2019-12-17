package binarytree

import "fmt"

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