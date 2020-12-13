package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/384/golang/solution"
)

func main() {
	/*

			["Solution","shuffle","reset","shuffle"]
		[[[1,2,3]],[],[],[]]

	*/

	var testName = "random-array"
	testLog := leet.NewTestLog(5)
	defer testLog.Render()

	idx := 1

	test := solution.Constructor([]int{1, 2, 3})
	got := test.Shuffle()
	sort.Ints(got)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testName, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testName, want))
	}
	idx++

	got = test.Reset()
	if !reflect.DeepEqual(got, want) {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testName, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testName, want))
	}
	idx++

	got = test.Shuffle()
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testName, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testName, want))
	}
	idx++
}
