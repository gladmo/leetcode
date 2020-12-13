package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/222/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	tests := []struct {
		name  string
		input *solution.TreeNode
		want  int
	}{
		{
			name:  "test-[1,2,3,4,5,6]",
			input: tree.CreateTree("[1,2,3,4,5,6]"),
			want:  6,
		},
		{
			name:  "test-[]",
			input: tree.CreateTree("[]"),
			want:  0,
		},
		{
			name:  "test-[1]",
			input: tree.CreateTree("[1]"),
			want:  1,
		},
		{
			name:  "test-[1,2,3,4]",
			input: tree.CreateTree("[1,2,3,4]"),
			want:  4,
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
