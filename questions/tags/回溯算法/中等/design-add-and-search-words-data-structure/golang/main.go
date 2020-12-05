package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/211/golang/solution"
)

func main() {

	testNamePre := "word-dictionary"
	testLog := leet.NewTestLog(6)
	defer testLog.Render()

	idx := 1

	trie := solution.Constructor()

	got := trie.Search("a")
	want := false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	trie.AddWord("bad")
	trie.AddWord("dad")
	trie.AddWord("mad")

	got = trie.Search("pad")
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	got = trie.Search("bad")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	got = trie.Search(".ad")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	got = trie.Search("b..")
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++
}
