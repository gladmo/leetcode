package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/784/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

			[4,2,7,1,3]
		5

	*/

	tests := []struct {
		name   string
		input1 *solution.TreeNode
		input2 int
		want   []string
	}{
		{
			name:   "test-[]",
			input1: tree.CreateTree("[]"),
			input2: 5,
			want: []string{
				"[5]",
			},
		},
		{
			name:   "test-[4,2,7,1,3]",
			input1: tree.CreateTree("[4,2,7,1,3]"),
			input2: 5,
			want: []string{
				"[4,2,7,1,3,5]",
				"[5,2,7,1,3,4]",
			},
		},
		{
			name:   "test-[40,20,60,10,30,50,70]",
			input1: tree.CreateTree("[40,20,60,10,30,50,70]"),
			input2: 25,
			want: []string{
				"[40,20,60,10,30,50,70,null,null,25]",
			},
		},
		{
			name:   "test-[4,2,7,1,3]",
			input1: tree.CreateTree("[4,2,7,1,3]"),
			input2: 5,
			want: []string{
				"[4,2,7,1,3,5]",
			},
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
		if !strings.Contains(strings.Join(test.want, ","), got.String()) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("result %v not in %v", got, test.want))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
