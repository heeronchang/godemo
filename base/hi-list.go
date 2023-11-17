package base

import (
	"container/list"
	"fmt"
)

func hiList() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	stack.PushFront(0)
	stack.PushFront(666)

	fmt.Printf("%v+\n", stack)

	for stack.Len() != 0 {
		fmt.Println(stack.Back().Value.(int))
		stack.Remove(stack.Back())
	}
}
