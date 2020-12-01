package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/816/golang/solution"
)

func main() {

	testLog := leet.NewTestLog(5)
	defer testLog.Render()

	idx := 1

	hash := solution.Constructor()
	hash.Add(1)
	hash.Add(2)
	got := hash.Contains(1)
	want := true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	got = hash.Contains(3)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	hash.Add(2)

	got = hash.Contains(2)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	hash.Remove(2)

	got = hash.Contains(2)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++
}
