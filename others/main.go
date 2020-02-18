package main

import (
    "fmt"
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

func main() {
    fmt.Println(checkValidSequence([]int{1,2,3,4,5}, []int{4, 5, 3, 2, 1}))
    fmt.Println(checkValidSequence([]int{1,2,3,4,5}, []int{4, 3, 5, 1, 2}))
    fmt.Println(checkValidSequence([]int{1,2,3,4,5}, []int{1, 2, 3, 4, 5}))
}