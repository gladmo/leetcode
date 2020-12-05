package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/208/golang/solution"
)

func main() {

	testLog := leet.NewTestLog(6)
	defer testLog.Render()

	idx := 1

	trie := solution.Constructor()

	got := trie.Search("a")
	want := false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	got = trie.StartsWith("app")
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("apple")

	got = trie.Search("apple")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	got = trie.Search("app")
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	got = trie.StartsWith("app")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("app")

	got = trie.Search("app")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++
}
