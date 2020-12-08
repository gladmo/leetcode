package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/236/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[3,5,1,6,2,0,8,null,null,7,4]
		5
		1

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 *solution.TreeNode
		input3 *solution.TreeNode
		want   *solution.TreeNode
	}{
		{
			name:   "test-[3,5,1,6,2,0,8,null,null,7,4]",
			input1: tree.CreateTree("[3,5,1,6,2,0,8,null,null,7,4]"),
			input2: &solution.TreeNode{Val: 5},
			input3: &solution.TreeNode{Val: 1},
			want:   &solution.TreeNode{Val: 3},
		},
		{
			name:   "test-[3,5,1,6,2,0,8,null,null,7,4]",
			input1: tree.CreateTree("[3,5,1,6,2,0,8,null,null,7,4]"),
			input2: &solution.TreeNode{Val: 5},
			input3: &solution.TreeNode{Val: 4},
			want:   &solution.TreeNode{Val: 5},
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

		got := solution.Export(test.input1, test.input2, test.input3)
		if !reflect.DeepEqual(test.want.Val, got.Val) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
