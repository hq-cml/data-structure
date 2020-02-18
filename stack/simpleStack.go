package stack

type SimpleStack []int

func (stack *SimpleStack) Len() int {
    return len(*stack)
}

func (stack *SimpleStack) IsEmpty() bool  {
    return len(*stack) == 0
}

func (stack *SimpleStack) Cap() int {
    return cap(*stack)
}

func (stack *SimpleStack) Push(value int)  {
    *stack = append(*stack, value)
}

//栈顶元素，但是栈不会变化
func (stack SimpleStack) Top() (int, bool)  {
    if len(stack) == 0 {
        return 0, false
    }
    return stack[len(stack) - 1], true
}

func (stack *SimpleStack) Pop() (int, bool)  {
    theStack := *stack
    if len(theStack) == 0 {
        return 0, false
    }
    value := theStack[len(theStack) - 1]
    *stack = theStack[:len(theStack) - 1]
    return value, true
}