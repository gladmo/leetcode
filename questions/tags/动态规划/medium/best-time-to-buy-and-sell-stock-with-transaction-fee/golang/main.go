package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/questions/tags/动态规划/medium/best-time-to-buy-and-sell-stock-with-transaction-fee/golang/solution"
	"github.com/gladmo/leetcode/leet"
)

func main() {
	/*
     
	[1,3,2,8,4,9]
2

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
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input)
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
