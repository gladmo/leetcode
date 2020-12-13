package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/380/golang/solution"
)

func main() {
	/*

			["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
		[[],[1],[2],[2],[],[1],[2],[]]

	*/
	testLog := leet.NewTestLog(5)
	defer testLog.Render()

	idx := 1

	set := solution.Constructor()
	got := set.Insert(1)
	want := true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Remove(2)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Insert(2)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got1 := set.GetRandom()
	fmt.Println("GetRandom:", got1)
	want1 := "[1,2]"
	if got1 != 1 && got1 != 2 {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want1), fmt.Sprintf("want: %v, got %v.", "[1,2]", want1))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want1))
	}
	idx++

	got = set.Remove(1)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Insert(2)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got1 = set.GetRandom()
	want1 = "[2]"
	if got1 != 2 {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want1), fmt.Sprintf("want: %v, got %v.", want1, got1))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want1))
	}
	idx++

	// empty
	got = set.Remove(2)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Remove(0)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Remove(0)
	want = false
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Insert(0)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got1 = set.GetRandom()
	want1 = "[0]"
	if got1 != 0 {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want1), fmt.Sprintf("want: %v, got %v.", want1, got1))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want1))
	}
	idx++

	got = set.Remove(0)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++

	got = set.Insert(0)
	want = true
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("randomized-set-%v", want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("randomized-set-%v", want))
	}
	idx++
}
