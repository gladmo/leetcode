package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/困难/239/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "test-[1,3,-1,-3,5,3,6,7]",
			input1: []int{1, 3, -1, -3, 5, 3, 6, 7},
			input2: 3,
			want:   []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:   "test-[1,2,3,4,5,6,7,8,9]",
			input1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			input2: 3,
			want:   []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:   "test-[-1,-2,-3,-4,-5,-6,-7,-8,-9]",
			input1: []int{-1, -2, -3, -4, -5, -6, -7, -8, -9},
			input2: 3,
			want:   []int{-1, -2, -3, -4, -5, -6, -7},
		},
		{
			name:   "test-[-1,-2,-3]",
			input1: []int{-1, -2, -3},
			input2: 3,
			want:   []int{-1},
		},
		{
			name:   "test-[1,2,3]",
			input1: []int{1, 2, 3},
			input2: 3,
			want:   []int{3},
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
