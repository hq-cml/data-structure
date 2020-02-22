package main

import (
    "fmt"
    "github.com/hq-cml/data-structure/stack"
)

func CheckSign(str string) bool {
    var mystack stack.Stack
    for _, t := range str {
        v := string(t)
        switch (v) {
        case ")":
            val, ok := mystack.Pop()
            if !ok {
                mystack.Push(v)
                continue
            }
            topVal, _ := val.(string)
            if topVal != "(" {
                return false
            }
        case "]":
            val, ok := mystack.Pop()
            if !ok {
                mystack.Push(v)
                continue
            }
            topVal, _ := val.(string)
            if topVal != "[" {
                return false
            }
        case "}":
            val, ok := mystack.Pop()
            if !ok {
                mystack.Push(v)
                continue
            }
            topVal, _ := val.(string)
            if topVal != "{" {
                return false
            }
        default:
            mystack.Push(v)
        }
    }

    if mystack.Len() > 0 {
       return false
    } else {
       return true
    }
}

func main() {
    fmt.Println(CheckSign("((("))
    fmt.Println(CheckSign("((()))"))
    fmt.Println(CheckSign("([()]}"))
    fmt.Println(CheckSign("{[()]}"))
}