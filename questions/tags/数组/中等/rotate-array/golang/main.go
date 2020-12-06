package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/189/golang/solution"
)

func main() {
	/*

			[1,2,3,4,5,6,7]
		3

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "test-[-1,-2]",
			input1: []int{-1, -2},
			input2: 1,
			want:   []int{-2, -1},
		},
		{
			name:   "test-[-1]",
			input1: []int{-1},
			input2: 2,
			want:   []int{-1},
		},
		{
			name:   "test-[1,2]",
			input1: []int{1, 2},
			input2: 3,
			want:   []int{2, 1},
		},
		{
			name:   "test-[1,2,3,4,5,6,7]",
			input1: []int{1, 2, 3, 4, 5, 6, 7},
			input2: 3,
			want:   []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:   "test-[-1,-100,3,99]",
			input1: []int{-1, -100, 3, 99},
			input2: 2,
			want:   []int{3, 99, -1, -100},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			// solution.Export(test.input1, test.input2)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		solution.Export(test.input1, test.input2)
		got := test.input1
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
