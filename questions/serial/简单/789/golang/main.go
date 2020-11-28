package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/questions/serial/简单/789/golang/solution"
	"github.com/gladmo/leetcode/leet"
)

func main() {
	/*
     
	["KthLargest","add","add","add","add","add"]
[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]

    */

	tests := []struct {
		name  string
		input [][]int
		want  bool
	}{
		{
			name: "test-[[1],[2],[3],[]]",
			input: [][]int{
				{1},
				{2},
				{3},
				{},
			},
			want: true,
		},
		{
			name: "test-[[1,3],[3,0,1],[2],[0]]",
			input: [][]int{
				{1, 3},
				{3, 0, 1},
				{2},
				{0},
			},
			want: false,
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
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
