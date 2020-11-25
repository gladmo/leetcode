package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/225/golang/solution"
)

func main() {
	/*

			["MyStack","push","push","top","pop","empty"]
		[[],[1],[2],[],[],[]]

	*/

	testLog := leet.NewTestLog(3)
	defer testLog.Render()

	stack := solution.Constructor()

	stack.Push(1)
	stack.Push(2)

	top := stack.Top()
	if top != 2 {
		testLog.Fail(1, "top", fmt.Sprintf("want %d, got %d.", 2, top))
	} else {
		testLog.Pass(1, "top")
	}

	pop := stack.Pop()
	if pop != 2 {
		testLog.Fail(2, "pop", fmt.Sprintf("want %d, got %d.", 2, pop))
	} else {
		testLog.Pass(2, "pop")
	}

	empty := stack.Empty()
	if empty != false {
		testLog.Fail(3, "empty", fmt.Sprintf("want %v, got %v.", false, empty))
	} else {
		testLog.Pass(3, "empty")
	}
}
