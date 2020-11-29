package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/110/golang/solution"
	"github.com/gladmo/leetcode/utils/tree"
)

func main() {
	/*

		[3,9,20,null,null,15,7]

	*/

	tests := []struct {
		name  string
		input *solution.TreeNode
		want  bool
	}{
		{
			name:  "test-[3,9,20,null,null,15,7]",
			input: tree.CreateTree("[3,9,20,null,null,15,7]"),
			want:  true,
		},
		{
			name:  "test-[1,2,2,3,3,null,null,4,4]",
			input: tree.CreateTree("[1,2,2,3,3,null,null,4,4]"),
			want:  false,
		},
		{
			name:  "test-[]",
			input: tree.CreateTree("[]"),
			want:  true,
		},
		{
			name:  "test-[1,2]",
			input: tree.CreateTree("[1,2]"),
			want:  true,
		},
		{
			name:  "test-[1,2,2,3,3]",
			input: tree.CreateTree("[1,2,2,3,3]"),
			want:  true,
		},
		{
			name:  "test-[1,null,2,null,3]",
			input: tree.CreateTree("[1,null,2,null,3]"),
			want:  false,
		},
		{
			name:  "test-[1,2,2,3,null,null,3,4,null,null,4]",
			input: tree.CreateTree("[1,2,2,3,null,null,3,4,null,null,4]"),
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
