package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/18/golang/solution"
	"github.com/gladmo/leetcode/utils/array"
)

func main() {
	/*

			[1,0,-1,0,-2,2]
		0

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   [][]int
	}{
		{
			name:   "test-[1,0,-1,0,-2,2]",
			input1: []int{1, 0, -1, 0, -2, 2},
			input2: 0,
			want: [][]int{
				{-1, 0, 0, 1},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
			},
		},
		{
			name:   "test-[1,0,-1,0,-2,2]",
			input1: []int{1, 0, -1, 0, -2, 2},
			input2: 2,
			want: [][]int{
				{-1, 0, 1, 2},
			},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2)
		if !array.OnlyOrderDifference(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
