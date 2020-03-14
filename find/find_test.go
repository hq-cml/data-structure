package main

import (
    "testing"
)
func TestFind(t *testing.T) {
    //异常输入
    a := &IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
    ret, ok := FindNum(a, -1)
    if ok {
        t.Error("Should be false!")
    }
    t.Log(ret)

    //常规
    a = &IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
    ret, ok = FindNum(a, 7)
    if !ok {
        t.Error("Should be true!")
    }
    t.Log(ret)

    //恰好第一个元素
    a = &IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
    ret, ok = FindNum(a, 0)
    if !ok {
        t.Error("Should be true!")
    }
    t.Log(ret)

    //恰好最后一元素
    a = &IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
    ret, ok = FindNum(a, 9)
    if !ok {
        t.Error("Should be true!")
    }
    t.Log(ret)

    //找不到
    a = &IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
    ret, ok = FindNum(a, 100)
    if ok {
        t.Error("Should be false!")
    }
    t.Log(ret)

    //元素全部相同
    a = &IntArr{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
    ret, ok = FindNum(a, 7)
    if !ok {
        t.Error("Should be true!")
    }
    t.Log(ret)

    //元素全部相同，且符合多个元素相加
    a = &IntArr{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
    ret, ok = FindNum(a, 14)
    if !ok {
        t.Error("Should be true!")
    }
    t.Log(ret)
}