package stack

type MinStackO1 struct {
    base Stack
    help Stack
}

func NewMinStackO1() *MinStackO1 {
    return &MinStackO1{
        base: Stack{},
        help: Stack{},
    }
}

func (stack *MinStackO1) Len() int {
    return stack.base.Len()
}

func (stack *MinStackO1) IsEmpty() bool  {
    return stack.base.IsEmpty()
}

func (stack *MinStackO1) Cap() int {
    return stack.base.Cap()
}

func (stack *MinStackO1) Push(value int)  {
    if !stack.help.IsEmpty() {
        v, _  := stack.help.Top()
        vi, _ := v.(int)
        if vi >= value {
            stack.help.Push(value)
        }
    } else {
        stack.help.Push(value)
    }
    stack.base.Push(value)
}

//栈顶元素，但是栈不会变化
func (stack MinStackO1) Top() (int, bool)  {
    v, ok := stack.base.Top()
    vi , _ := v.(int)
    return vi, ok
}

func (stack *MinStackO1) Pop() (int, bool)  {
    if !stack.help.IsEmpty() {
        v1, _  := stack.help.Top()
        v2, _  := stack.base.Top()
        vi1, _ := v1.(int)
        vi2, _ := v2.(int)
        if vi2 <= vi1 {
            stack.help.Pop()
        }
    }

    v, ok := stack.base.Pop()
    vi , _ := v.(int)
    return vi, ok
}

func (stack *MinStackO1) Min() (int, bool)  {
    if !stack.help.IsEmpty() {
        v1, _  := stack.help.Top()
        vi1, _ := v1.(int)
        return vi1, true
    }
    return 0, false
}