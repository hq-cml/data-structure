package stack

type Stack []interface {}

func (stack *Stack) Len() int {
    return len(*stack)
}

func (stack *Stack) IsEmpty() bool  {
    return len(*stack) == 0
}

func (stack *Stack) Cap() int {
    return cap(*stack)
}

func (stack *Stack) Push(value interface{})  {
    *stack = append(*stack, value)
}

//栈顶元素，但是栈不会变化
func (stack Stack) Top() (interface{}, bool)  {
    if len(stack) == 0 {
        return nil, false
    }
    return stack[len(stack) - 1], true
}

func (stack *Stack) Pop() (interface{}, bool)  {
    theStack := *stack
    if len(theStack) == 0 {
        return nil, false
    }
    value := theStack[len(theStack) - 1]
    *stack = theStack[:len(theStack) - 1]
    return value, true
}