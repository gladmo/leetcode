package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/677/golang/solution"
)

func main() {

	testLog := leet.NewTestLog(7)
	defer testLog.Render()

	idx := 1

	trie := solution.Constructor()

	got := trie.Sum("ap")
	want := 0
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("apple", 3)

	got = trie.Sum("ap")
	fmt.Println("sum ap", got)
	want = 3
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("app", 2)

	got = trie.Sum("ap")
	want = 5
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	got = trie.Sum("a")
	want = 5
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("app", 3)
	got = trie.Sum("a")
	want = 6
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("apple", 1)

	got = trie.Sum("a")
	want = 4
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++

	trie.Insert("apples", 2)
	got = trie.Sum("a")
	want = 6
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("hash-table-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("hash-table-%v", want))
	}
	idx++
}
