package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/1242/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 [][]int
		input2 int
		want   [][]int
	}{
		{
			name: "test-[[1,2,3],[4,5,6],[7,8,9]]-1",
			input1: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
			},
			input2: 1,
			want: [][]int{
				{12, 21, 16}, {27, 45, 33}, {24, 39, 28},
			},
		},
		{
			name: "test-[[1,2,3],[4,5,6],[7,8,9]]-2",
			input1: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
			},
			input2: 2,
			want: [][]int{
				{45, 45, 45}, {45, 45, 45}, {45, 45, 45},
			},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input1, test.input2)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
