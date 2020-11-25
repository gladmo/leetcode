package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/232/golang/solution"
)

func main() {
	/*

			["MyQueue","push","push","peek","pop","empty"]
		[[],[1],[2],[],[],[]]

	*/

	testLog := leet.NewTestLog(3)
	defer testLog.Render()

	queue := solution.Constructor()
	queue.Push(1)
	queue.Push(2)

	peek := queue.Peek()
	if peek != 1 {
		testLog.Fail(1, "peek", fmt.Sprintf("want %d, got %d.", 1, peek))
	} else {
		testLog.Pass(1, "peek")
	}

	pop := queue.Pop()
	if pop != 1 {
		testLog.Fail(2, "pop", fmt.Sprintf("want %d, got %d.", 1, pop))
	} else {
		testLog.Pass(2, "pop")
	}

	empty := queue.Empty()
	if empty != false {
		testLog.Fail(3, "empty", fmt.Sprintf("want %v, got %v.", false, empty))
	} else {
		testLog.Pass(3, "empty")
	}
}
