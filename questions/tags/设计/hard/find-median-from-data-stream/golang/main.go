package main

import (
	"fmt"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/hard/295/golang/solution"
)

func main() {

	testNamePre := "median-finder"
	testLog := leet.NewTestLog(3)
	defer testLog.Render()

	idx := 1

	median := solution.Constructor()
	median.AddNum(1)
	median.AddNum(2)

	got := median.FindMedian()
	want := 1.5
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	median.AddNum(3)

	got = median.FindMedian()
	want = 2
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++

	median.AddNum(3)
	median.AddNum(3)
	median.AddNum(4)
	median.AddNum(5)
	median.AddNum(5)
	median.AddNum(5)
	median.AddNum(6)
	median.AddNum(1)
	median.AddNum(0)
	median.AddNum(0)
	median.AddNum(0)
	median.AddNum(6)
	median.AddNum(6)
	median.AddNum(6)

	fmt.Println(median.Data)
	got = median.FindMedian()
	want = 3
	if got != want {
		testLog.Fail(idx, fmt.Sprintf("%s-%v", testNamePre, want), fmt.Sprintf("want: %v, got %v.", want, got))
	} else {
		testLog.Pass(idx, fmt.Sprintf("%s-%v", testNamePre, want))
	}
	idx++
}
