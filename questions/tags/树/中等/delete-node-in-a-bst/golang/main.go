package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/450/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[5,3,6,2,4,null,7]
		3

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 int
		want   []string
	}{
		{
			name:   "test-[5,2,6,1,4,null,7,null,null,3]",
			input1: tree.CreateTree("[5,2,6,1,4,null,7,null,null,3]"),
			input2: 5,
			want: []string{
				"[6,3,7,2,4]",
				"[6,2,7,1,4,null,null,null,null,3]",
			},
		},
		{
			name:   "test-[5,2,6,1,4,null,7,null,null,3]",
			input1: tree.CreateTree("[5,2,6,1,4,null,7,null,null,3]"),
			input2: 7,
			want: []string{
				"[5,2,6,1,4,null,null,null,null,3]",
			},
		},
		{
			name:   "test-[5,2,6,1,4,null,7,null,null,3]",
			input1: tree.CreateTree("[5,2,6,1,4,null,7,null,null,3]"),
			input2: 6,
			want: []string{
				"[5,2,7,1,4,null,null,null,null,3]",
			},
		},
		{
			name:   "test-[5,2,6,1,4,null,7,null,null,3]",
			input1: tree.CreateTree("[5,2,6,1,4,null,7,null,null,3]"),
			input2: 2,
			want: []string{
				"[5,3,6,1,4,null,7]",
			},
		},
		{
			name:   "test-[5,3,6,2,4,null,7]",
			input1: tree.CreateTree("[5,3,6,2,4,null,7]"),
			input2: 3,
			want: []string{
				"[5,4,6,2,null,null,7]",
				"[5,2,6,null,4,null,7]",
			},
		},
		{
			name:   "test-[5,3,6,2,4,null,7]",
			input1: tree.CreateTree("[5,3,6,2,4,null,7]"),
			input2: 0,
			want: []string{
				"[5,3,6,2,4,null,7]",
			},
		},
		{
			name:   "test-[0]",
			input1: tree.CreateTree("[0]"),
			input2: 0,
			want: []string{
				"[]",
			},
		},
		{
			name:   "test-[]",
			input1: tree.CreateTree("[]"),
			input2: 0,
			want: []string{
				"[]",
			},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {

			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2)
		if !strings.Contains(strings.Join(test.want, ","), got.String()) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
