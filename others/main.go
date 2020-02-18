package main

import (
    "fmt"
    "github.com/hq-cml/data-structure/binarytree"
    stack2 "github.com/hq-cml/data-structure/stack"
)

func checkValidSequence(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    length := len(a)
    if length == 0 {
        return true
    }

    i, j := 0, 0
    stack := stack2.SimpleStack{} //辅助栈

    for i < length {

        for a[i] != b[j] && i < length {
            stack.Push(a[i])
            i ++
        }
        if i == length {
            if !stack.IsEmpty() {
                return false
            } else {
                return true
            }
        }
        if a[i] == b[j] {
            stack.Push(a[i])
            i ++
        }
        if !stack.IsEmpty() {
            for {
                v, ok := stack.Top()
                if !ok {
                    break
                }
                if v == b[j] {
                    j ++
                    stack.Pop()
                } else {
                    break
                }
            }
        }
    }
    if !stack.IsEmpty() {
        return false
    } else {
        return true
    }
}

//按行打印二叉树
func PrintTree1(root *binarytree.TreeNode) {
    if root == nil {
        return
    }
    list := binarytree.NewSimpleList()
    list.Put(root)

    for list.Len() > 0 {
        node := list.Get()
        fmt.Print(node.Data, " ")
        if node.Left != nil {
            list.Put(node.Left)
        }
        if node.Right != nil {
            list.Put(node.Right)
        }
    }
}

//按行打印二叉树
func PrintTree2(root *binarytree.TreeNode) {
    if root == nil {
        return
    }
    list := binarytree.NewSimpleList()
    list.Put(root)
    currentLevelCnt := 1;
    nextLevelCnt := 0;
    for list.Len() > 0 {
        node := list.Get()
        fmt.Print(node.Data, " ")
        currentLevelCnt --
        if node.Left != nil {
            list.Put(node.Left)
            nextLevelCnt ++
        }
        if node.Right != nil {
            list.Put(node.Right)
            nextLevelCnt ++
        }
        if currentLevelCnt == 0 {
            currentLevelCnt = nextLevelCnt
            nextLevelCnt = 0
            fmt.Println()
        }
    }
}

func main() {
    tree := binarytree.NewTree(1,
        binarytree.NewTree(2,
            binarytree.NewTree(4, nil, nil),
            binarytree.NewTree(5, nil, nil),
        ),
        binarytree.NewTree(3,
            binarytree.NewTree(6, nil, nil),
            binarytree.NewTree(7, nil, nil),
        ),
    )
    PrintTree2(tree)
}