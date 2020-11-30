package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/162/golang/solution"
)

func main() {
	/*

		[1,2,3,1]

	*/

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test-[1,2,3,1]",
			input: []int{1, 2, 3, 1},
			want:  []int{2},
		},
		{
			name:  "test-[1,2,1,3,5,6,4]",
			input: []int{1, 2, 1, 3, 5, 6, 4},
			want:  []int{1, 5},
		},
		{
			name:  "test-[1]",
			input: []int{1},
			want:  []int{0},
		},
		{
			name:  "test-[2,1]",
			input: []int{2, 1},
			want:  []int{0},
		},
		{
			name:  "test-[1,2]",
			input: []int{1, 2},
			want:  []int{1},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input)
		success := false
		sort.Search(len(test.want), func(i int) bool {
			if test.want[i] == got {
				success = true
			}
			return success
		})
		if !success {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
