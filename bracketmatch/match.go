package main

import (
    "fmt"
    "github.com/hq-cml/data-structure/stack"
)

func CheckSign(str string) bool {
    var mystack stack.Stack
    for _, t := range str {
        v := string(t)
        val, ok := mystack.Top()
        if !ok {
            mystack.Push(v)
            continue
        }
        topVal, _ := val.(string)

        switch string(v) {
        case ")":
            if topVal == "(" {
                mystack.Pop()
            } else if topVal == "[" || topVal == "{"{
                fmt.Println("Fuck 1")
                return false
            } else {
                mystack.Push(v)
            }
        case "]":
            if topVal == "[" {
                mystack.Pop()
            } else if topVal == "(" || topVal == "{"{
                fmt.Println("Fuck 2")
                return false
            } else {
                mystack.Push(v)
            }
        case "}":
            if topVal == "{" {
                mystack.Pop()
            } else if topVal == "[" || topVal == "("{
                fmt.Println("Fuck 3")
                return false
            } else {
                mystack.Push(v)
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
