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

type SimpleList struct {
    list []*TreeNode
}

func NewSimpleList() *SimpleList{
    return &SimpleList{list:[]*TreeNode{}}
}

func (s *SimpleList)Len() int{
    return len(s.list)
}

func (s *SimpleList)Put(node *TreeNode) {
    s.list = append(s.list, node)
}

func (s *SimpleList)Get() *TreeNode {
    if s.Len() > 0 {
        t := s.list[0]
        s.list = s.list[1:]
        return t
    }
    return nil
}

//路径和求值
func(tree *TreeNode) FindPathSum(num int) {
    s := []int{}
    findPathSum(tree, 22, s)
}

func findPathSum(root *TreeNode, num int, s []int) {
    if root == nil {
        return
    }
    v, _ := root.Data.(int)
    s = append(s, v)
    sum := 0
    for _,v := range s {
        sum += v
    }
    if sum == num && root.Left == nil && root.Right == nil{
        fmt.Println(s)
    }
    findPathSum(root.Left, num, s)
    findPathSum(root.Right, num, s)
}

//求指定值的路径
func(tree *TreeNode) FindPath(num int, path []int) ([]int, bool) {
    if tree == nil {
        return path, false
    }
    data, _ := tree.Data.(int)
    path = append(path, data)
    //if data == num && tree.Left == nil && tree.Right == nil {
    if data == num {
        return path, true
    }

    ret, ok := tree.Left.FindPath(num, path)
    if ok {
        return ret, true
    }
    ret, ok = tree.Right.FindPath(num, path)
    if ok {
        return ret, true
    }

    return path, false
}

//求公共节点
func(tree *TreeNode) FindCommonParent(num1, num2 int) (int, bool) {
    path1 := []int{}
    path2 := []int{}

    path1, ok1 := tree.FindPath(num1, path1)
    path2, ok2 := tree.FindPath(num2, path2)

    if !ok1 || !ok2 {
        return 0, false
    }

    for k, v := range path1 {
        if k > len(path2)-1 {
            //paht1长
            return path1[k-1], true;
        }

        if v != path2[k] {
            return path1[k-1], true;
        }
    }
    return 0, false
}

//判断给定树是否是子树
func(tree *TreeNode) CheckSubTree(tree2 *TreeNode) bool {
    if tree == nil && tree2 == nil {
        return true
    }

    if tree2 == nil {
        return true
    }

    ret := false
    if tree.Data == tree2.Data {
        ret = checkSubTree(tree, tree2)
    }

    if !ret {
        ret = checkSubTree(tree.Left, tree2)
    }

    if !ret {
        ret = checkSubTree(tree.Right, tree2)
    }

    return ret
}

func checkSubTree(tree1, tree2 *TreeNode) bool {
    if tree2 == nil {
        return true
    }

    if tree1 == nil {
        return false
    }

    if tree1.Data != tree2.Data {
        return false
    }

    return checkSubTree(tree1.Left, tree2.Left) &&
             checkSubTree(tree1.Right, tree2.Right)
}

//二叉树镜像
func(tree *TreeNode) MirrorTree()  {
    if tree == nil {
        return
    }
    if tree.Right == nil && tree.Left == nil {
        return
    }

    //左右互换
    tree.Left, tree.Right = tree.Right, tree.Left

    //左右各自递归变更
    tree.Left.MirrorTree()
    tree.Right.MirrorTree()

    return
}