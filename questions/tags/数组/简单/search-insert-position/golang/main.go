package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/35/golang/solution"
)

func main() {
	/*

			[1,3,5,6]
		5

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   int
	}{
		{
			name:   "test-[1,3,5,6]-5",
			input1: []int{1, 3, 5, 6},
			input2: 5,
			want:   2,
		},
		{
			name:   "test-[1,3,5,6]-7",
			input1: []int{1, 3, 5, 6},
			input2: 7,
			want:   4,
		},
		{
			name:   "test-[1,3,5,6]-0",
			input1: []int{1, 3, 5, 6},
			input2: 0,
			want:   0,
		},
		{
			name:   "test-[1,3,5,6]-2",
			input1: []int{1, 3, 5, 6},
			input2: 2,
			want:   1,
		},
		{
			name:   "test-[1]-2",
			input1: []int{1},
			input2: 2,
			want:   1,
		},
		{
			name:   "test-[1,4,6,7,8,9]-6",
			input1: []int{1, 4, 6, 7, 8, 9},
			input2: 6,
			want:   1,
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
