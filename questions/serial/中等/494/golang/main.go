package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/494/golang/solution"
)

func main() {
	/*

			[1,1,1,1,1]
		3

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   int
	}{
		{
			name:   "test-[1, 1, 1, 1, 1]",
			input1: []int{1, 1, 1, 1, 1},
			input2: 3,
			want:   5,
		},
		{
			name:   "test-[1, 0]",
			input1: []int{1, 0},
			input2: 1,
			want:   2,
		},
		{
			name:   "test-[]",
			input1: []int{},
			input2: 0,
			want:   0,
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

		result := solution.Export(test.input1, test.input2)
		if result != test.want {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("expect %d, got %d.", test.want, result))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
