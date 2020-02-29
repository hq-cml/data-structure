package main

import (
    "errors"
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

//寻找最长不重复子串（滑动窗口）
func FindMaxSubString2(str string) (int, int) {
    var maxCnt = 0;
    var index = 0;
    length := len(str)
    i, j := 0, 0
    m := map[byte]struct{}{}
    for i < length && j < length {
        c := str[i]
        _, ok := m[c]
        if !ok {
            //窗口右边界右移
            i++
            m[c] = struct{}{}
            if maxCnt < (i-j) {
                maxCnt = (i - j)
                index = j
            }
        } else {
            //窗口左边界右移
            delete(m, str[j])
            j++
        }
    }

    return index, maxCnt
}

func FindMin(arr []int) (int, error) {
    if arr == nil {
        return 0, errors.New("Invalid input")
    }

    length := len(arr)
    if length == 1 {
        return arr[0], nil
    }
    if length == 2 {
        if arr[0] > arr[1] {
            return arr[1], nil
        }else {
            return arr[0], nil
        }
    }

    if arr[0] <= arr[length-1] {
        return arr[0], nil
    }

    mid := length/2

    if (arr[0] <= arr[mid] && arr[mid+1] <= arr[length-1]) {
        return arr[mid+1], nil
    }

    if (arr[0] > arr[mid]) {
        return FindMin(arr[:mid+1])
    }

    if (arr[mid+1] > arr[length-1]) {
        return FindMin(arr[mid+1:])
    }
    return 0, errors.New("Somthig wrong")
}

type A struct {
    a int
}
func main() {
    fmt.Println(FindMaxSubString2("aaa"))
}