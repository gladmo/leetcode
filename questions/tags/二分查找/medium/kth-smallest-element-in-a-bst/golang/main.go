package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/230/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[3,1,4,null,2]
		1

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 int
		want   int
	}{
		{
			name:   "test-[3,1,4,null,2]",
			input1: tree.CreateTree("[3,1,4,null,2]"),
			input2: 1,
			want:   1,
		},
		{
			name:   "test-[5,3,6,2,4,null,null,1]",
			input1: tree.CreateTree("[5,3,6,2,4,null,null,1]"),
			input2: 3,
			want:   3,
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
