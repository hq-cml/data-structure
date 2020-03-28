package main

import (
    "fmt"
    "github.com/hq-cml/data-structure/stack"
)
type MaxWindowArr struct {
    Arr []int
    Help *MaxList
}

type MaxList struct {
    St1 *stack.MaxStackO1
    St2 *stack.MaxStackO1
}
func NewMaxList() *MaxList {
    return &MaxList{
        St1: stack.NewMaxStackO1(),
        St2: stack.NewMaxStackO1(),
    }
}

func (ml *MaxList) Insert(value int)  {
    ml.St1.Push(value)
}
func (ml *MaxList) Del()  {
    if !ml.St2.IsEmpty() {
        ml.St2.Pop()
        return
    }
    if  !ml.St1.IsEmpty() {
        for !ml.St1.IsEmpty() {
            v, _ := ml.St1.Pop()
            ml.St2.Push(v)
        }
        ml.St2.Pop()
        return
    }
    return
}
func (ml *MaxList) Max() (int, bool){
    mx1, ok1 := ml.St1.Max()
    mx2, ok2 := ml.St2.Max()

    if !ok1 && !ok2 {
        return 0, false
    }

    if ok1 && ok2 {
        if mx1 > mx2 {
            return mx1, true
        } else {
            return mx2, true
        }
    }

    if ok1 {
        return mx1, true
    } else {
        return mx2, true
    }
}

func (wa *MaxWindowArr) FindWindowMax(arr []int, windowLen int) ([]int, bool)  {
    wa = &MaxWindowArr{
        Arr:  arr,
        Help: NewMaxList(),
    }

    if arr == nil {
        return nil, false
    }

    if len(arr) <= windowLen {
        for _, v := range wa.Arr {
            wa.Help.Insert(v)
        }
        max, ok := wa.Help.Max()
        if ok {
            return []int{max}, true
        }
        return nil, false
    }

    ret := []int{}
    for k, v := range wa.Arr {
        if k <= (windowLen-1) {
            wa.Help.Insert(v)
            continue
        }
        max, ok := wa.Help.Max()
        if ok {
            ret = append(ret, max)
        }
        wa.Help.Del()
        wa.Help.Insert(wa.Arr[k])
    }

    max, ok := wa.Help.Max()
    if ok {
        ret = append(ret, max)
    }
    return ret, true
}

func Add(i int) int {
    if i == 0 {
        return 0
    }
    return i + Add(i-1)
}

func main() {
    fmt.Println(Add(3))
}