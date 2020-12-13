package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/871/golang/solution"
)

func main() {
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
