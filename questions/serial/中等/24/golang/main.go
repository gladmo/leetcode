package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/24/golang/solution"
	"github.com/gladmo/leetcode/utils/list"
)

func main() {

	tests := []struct {
		name  string
		input *solution.ListNode
		want  string
	}{
		{
			name:  "test-[1,2,3,4]",
			input: list.CreateNode("[1,2,3,4]"),
			want:  "[2,1,4,3]",
		},
		{
			name:  "test-[]",
			input: list.CreateNode("[]"),
			want:  "[]",
		},
		{
			name:  "test-[1]",
			input: list.CreateNode("[1]"),
			want:  "[1]",
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input)
			solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input)
		if !reflect.DeepEqual(test.want, got.String()) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
