package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/103/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {

	tests := []struct {
		name  string
		input *solution.TreeNode
		want  [][]int
	}{
		{
			name:  "test-[3,9,20,null,null,15,7]",
			input: tree.CreateTree("[3,9,20,null,null,15,7]"),
			want: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			name:  "test-[1,2,3,4,null,null,5]",
			input: tree.CreateTree("[1,2,3,4,null,null,5]"),
			want: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
		{
			name:  "test-...",
			input: tree.CreateTree("[5,0,-4,-1,-6,-9,null,7,null,1,3,null,0,null,9,null,null,6,0,null,-7,null,null,null,null,null,null,-4,null,1,null,null,-4]"),
			want: [][]int{
				{5},
				{-4, 0},
				{-1, -6, -9},
				{0, 3, 1, 7},
				{9, 6, 0, -7},
				{-4},
				{1},
				{-4},
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
