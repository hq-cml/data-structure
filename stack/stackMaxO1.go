package stack

type MaxStackO1 struct {
    base Stack
    help Stack
}

func NewMaxStackO1() *MaxStackO1 {
    return &MaxStackO1{
        base: Stack{},
        help: Stack{},
    }
}

func (stack *MaxStackO1) Len() int {
    return stack.base.Len()
}

func (stack *MaxStackO1) IsEmpty() bool  {
    return stack.base.IsEmpty()
}

func (stack *MaxStackO1) Cap() int {
    return stack.base.Cap()
}

func (stack *MaxStackO1) Push(value int)  {
    if !stack.help.IsEmpty() {
        v, _  := stack.help.Top()
        vi, _ := v.(int)
        if vi <= value {
            stack.help.Push(value)
        }
    } else {
        stack.help.Push(value)
    }
    stack.base.Push(value)
}

//栈顶元素，但是栈不会变化
func (stack MaxStackO1) Top() (int, bool)  {
    v, ok := stack.base.Top()
    vi , _ := v.(int)
    return vi, ok
}

func (stack *MaxStackO1) Pop() (int, bool)  {
    if !stack.help.IsEmpty() {
        v1, _  := stack.help.Top()
        v2, _  := stack.base.Top()
        vi1, _ := v1.(int)
        vi2, _ := v2.(int)
        if vi2 >= vi1 {
            stack.help.Pop()
        }
    }

    v, ok := stack.base.Pop()
    vi , _ := v.(int)
    return vi, ok
}

func (stack *MaxStackO1) Max() (int, bool)  {
    if !stack.help.IsEmpty() {
        v1, _  := stack.help.Top()
        vi1, _ := v1.(int)
        return vi1, true
    }
    return 0, false
}