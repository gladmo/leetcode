package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/173/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]
		[[[7,3,15,null,null,9,20]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]

	*/

	testLog := leet.NewTestLog(9)
	defer testLog.Render()

	idx := 1

	bst := solution.Constructor(tree.CreateTree("[7,3,15,null,null,9,20]"))
	got := bst.Next()
	want := 3
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", want))
	}
	idx++

	got = bst.Next()
	want = 7
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", want))
	}
	idx++

	gotT := bst.HasNext()
	if gotT != true {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", true), fmt.Sprintf("want: %v, got %v.", true, gotT))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", true))
	}
	idx++

	got = bst.Next()
	want = 9
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", want))
	}
	idx++

	gotT = bst.HasNext()
	if gotT != true {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", true), fmt.Sprintf("want: %v, got %v.", true, gotT))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", true))
	}
	idx++

	got = bst.Next()
	want = 15
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", want))
	}
	idx++

	gotT = bst.HasNext()
	if gotT != true {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", true), fmt.Sprintf("want: %v, got %v.", true, gotT))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", true))
	}
	idx++

	got = bst.Next()
	want = 20
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", want))
	}
	idx++

	gotT = bst.HasNext()
	if gotT != false {
		testLog.Fail(idx, fmt.Sprintf("bst-%v", false), fmt.Sprintf("want: %v, got %v.", false, gotT))
	} else {
		testLog.Pass(idx, fmt.Sprintf("bst-%v", false))
	}
	idx++
}
