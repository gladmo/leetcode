package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/33/golang/solution"
)

func main() {
	/*

			[4,5,6,7,0,1,2]
		0

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   int
	}{
		{
			name:   "test-[4,5,6,7,0,1,2]-0",
			input1: []int{4, 5, 6, 7, 0, 1, 2},
			input2: 0,
			want:   4,
		},
		{
			name:   "test-[4,5,6,7,0,1,2]-3",
			input1: []int{4, 5, 6, 7, 0, 1, 2},
			input2: 3,
			want:   -1,
		},
		{
			name:   "test-[4,5,6,7,0,1,2]-3",
			input1: []int{4, 5, 6, 7, 0, 1, 2},
			input2: 7,
			want:   3,
		},
		{
			name:   "test-[1]-0",
			input1: []int{1},
			input2: 0,
			want:   -1,
		},
		{
			name:   "test-[2,1]-1",
			input1: []int{2, 1},
			input2: 1,
			want:   1,
		},
		{
			name:   "test-[2,1]-0",
			input1: []int{2, 1},
			input2: 0,
			want:   -1,
		},
		{
			name:   "test-[5,1,3]-5",
			input1: []int{5, 1, 3},
			input2: 5,
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

		got := solution.Export(test.input1, test.input2)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
