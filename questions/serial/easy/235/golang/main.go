package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/235/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[6,2,8,0,4,7,9,null,null,3,5]
		2
		8

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 *solution.TreeNode
		input3 *solution.TreeNode
		want   int
	}{
		{
			name:   "test-[6,2,8,0,4,7,9,null,null,3,5]",
			input1: tree.CreateTree("[6,2,8,0,4,7,9,null,null,3,5]"),
			input2: tree.CreateTree("[2]"),
			input3: tree.CreateTree("[8]"),
			want:   6,
		},
		{
			name:   "test-[6,2,8,0,4,7,9,null,null,3,5]",
			input1: tree.CreateTree("[6,2,8,0,4,7,9,null,null,3,5]"),
			input2: tree.CreateTree("[2]"),
			input3: tree.CreateTree("[4]"),
			want:   2,
		},
		{
			name:   "test-[6,2,8,0,4,7,9,null,null,3,5]",
			input1: tree.CreateTree("[6,2,8,0,4,7,9,null,null,3,5]"),
			input2: tree.CreateTree("[7]"),
			input3: tree.CreateTree("[9]"),
			want:   8,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2, test.input3)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2, test.input3).Val
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
