package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/783/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[4,2,7,1,3]
		2

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 int
		want   string
	}{
		{
			name:   "test-[4,2,7,1,3]",
			input1: tree.CreateTree("[4,2,7,1,3]"),
			input2: 2,
			want:   "[2,1,3]",
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
		if !reflect.DeepEqual(test.want, got.String()) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got.String()))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
