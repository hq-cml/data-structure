package main

import "sync"

type Singtleton struct{}

var ins *Singtleton
var mut sync.Mutex
var once sync.Once

func GetInstance1() *Singtleton {
    mut.Lock()
    defer mut.Unlock()

    if ins == nil {
        ins = &Singtleton{}
    }
    return ins
}

func GetInstance2() *Singtleton {
    if ins == nil {
        mut.Lock()
        defer mut.Unlock()

        if ins == nil {
            ins = &Singtleton{}
        }
    }
    return ins
}

func GetInstance3() *Singtleton {
    once.Do(func(){
        ins = &Singtleton{}
    })
    return ins
}