package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/789/golang/solution"
)

func main() {

	testLog := leet.NewTestLog(5)
	defer testLog.Render()

	idx := 1

	kth := solution.Constructor(3, []int{4, 5, 8, 2})
	got := kth.Add(3)
	want := 4
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("kth-largest-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("kth-largest-%v", want))
	}
	idx++

	got = kth.Add(5)
	want = 5
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("kth-largest-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("kth-largest-%v", want))
	}
	idx++

	got = kth.Add(10)
	want = 5
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("kth-largest-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("kth-largest-%v", want))
	}
	idx++

	got = kth.Add(9)
	want = 8
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("kth-largest-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("kth-largest-%v", want))
	}
	idx++

	got = kth.Add(4)
	want = 8
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("kth-largest-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("kth-largest-%v", want))
	}
	idx++

	// test 2
	kth = solution.Constructor(1, []int{})
	got = kth.Add(-3)
	got = kth.Add(-2)
	got = kth.Add(-4)
	got = kth.Add(0)
	got = kth.Add(4)

}
