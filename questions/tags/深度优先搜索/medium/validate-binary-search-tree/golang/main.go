package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/98/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {

	tests := []struct {
		name  string
		input *solution.TreeNode
		want  bool
	}{
		{
			name:  "test-[2,1,3]",
			input: tree.CreateTree("[2,1,3]"),
			want:  true,
		},
		{
			name:  "test-[5,1,4,null,null,3,6]",
			input: tree.CreateTree("[5,1,4,null,null,3,6]"),
			want:  false,
		},
		{
			name:  "test-[1,1]",
			input: tree.CreateTree("[1,1]"),
			want:  false,
		},
		{
			name:  "test-[1,null,1]",
			input: tree.CreateTree("[1,null,1]"),
			want:  false,
		},
		{
			name:  "test-[10,5,15,null,null,6,20]",
			input: tree.CreateTree("[10,5,15,null,null,6,20]"),
			want:  false,
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
