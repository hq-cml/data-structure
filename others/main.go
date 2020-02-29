package main

import (
    "fmt"
    "github.com/hq-cml/data-structure/binarytree"
    stack2 "github.com/hq-cml/data-structure/stack"
    "reflect"
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

//寻找最长不重复子串
func FindMaxSubString(str string) (int, int) {
    var maxCnt = 0;
    var index = 0;
    for idx:=0; idx < len(str); idx ++ {
        cnt := 0;
        m := map[string]struct{}{}
        for _, c := range str[idx:] {
            _, ok := m[string(c)]
            if !ok {
                m[string(c)] = struct{}{}
                cnt ++
            } else {
                if cnt > maxCnt {
                    maxCnt = cnt
                    index = idx
                    break
                }
            }
        }
        if cnt > maxCnt {
            maxCnt = cnt
            index = idx
        }
    }

    return index, maxCnt
}

type A struct {
    a int
}
func main() {
    //fmt.Println(FindMaxSubString("aabc"));
    str := "abc我d"
    for k, v := range str {
        fmt.Println(k, v, reflect.TypeOf(v))
    }
    fmt.Println()
    str2 := []rune(str)
    for k, v := range str2 {
        fmt.Println(k, v, reflect.TypeOf(v))
    }
}