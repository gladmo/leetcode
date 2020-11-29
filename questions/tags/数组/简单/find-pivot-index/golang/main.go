package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/724/golang/solution"
)

func main() {
	/*

		[1,7,3,6,5,6]

	*/

	tests := []struct {
		name   string
		input1 []int
		want   int
	}{
		{
			name:   "test-[1, 7, 3, 6, 5, 6]",
			input1: []int{1, 7, 3, 6, 5, 6},
			want:   3,
		},
		{
			name:   "test-[1, 2, 3]",
			input1: []int{1, 2, 3},
			want:   -1,
		},
		{
			name:   "test-[1, 7, 3, 7, 6, 5, 6, 7]",
			input1: []int{1, 7, 3, 7, 6, 5, 6, 7},
			want:   4,
		},
		{
			name:   "test-[]",
			input1: []int{},
			want:   -1,
		},
		{
			name:   "test-[-1,-1,-1,-1,-1,-1]",
			input1: []int{-1, -1, -1, -1, -1, -1},
			want:   -1,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
