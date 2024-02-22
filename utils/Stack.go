package utils

import (
	"sync"
)

// Item interface to store any data type in stack
// type Item interface{}

// Stack struct which contains a list of Items
type Stack struct {
    items []int
    mutex sync.Mutex
}

// NewEmptyStack() returns a new instance of Stack with zero elements
func NewEmptyStack() *Stack {
    return &Stack{
        items: nil,
    }
}

// NewStack() returns a new instance of Stack with list of specified elements
func NewStack(items []int) *Stack {
    return &Stack{
        items: items,
    }
}

// Push() adds new item to top of existing/empty stack
func (stack *Stack) Push(item int) {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    stack.items = append(stack.items, item)
}

// Pop() removes most recent item(top) from stack
func (stack *Stack) Pop() (int, bool) {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    if len(stack.items) == 0 {
        return 0, false;
    }

    lastItem := stack.items[len(stack.items)-1]
    stack.items = stack.items[:len(stack.items)-1]

    return lastItem, true;
}

// IsEmpty() returns whether the stack is empty or not (boolean result)
func (stack *Stack) IsEmpty() bool {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    return len(stack.items) == 0
}

// Top() returns the last inserted item in stack without removing it.
func (stack *Stack) Top() (int, bool) {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    if len(stack.items) == 0 {
        return 0, false;
    }

    return stack.items[len(stack.items)-1], true;
}

// Size() returns the number of items currently in the stack
func (stack *Stack) Size() int {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    return len(stack.items)
}

// Clear() removes all items from the stack
func (stack *Stack) Clear() {
    stack.mutex.Lock()
    defer stack.mutex.Unlock()

    stack.items = nil
}