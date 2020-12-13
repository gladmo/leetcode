package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/21/golang/solution"
	"github.com/gladmo/leetcode/utils/list"
)

func main() {

	tests := []struct {
		name           string
		input1, input2 *solution.ListNode
		want           string
	}{
		{
			name:   "test-[1,2,4]+[1,3,4]",
			input1: list.CreateNode("[1,2,4]"),
			input2: list.CreateNode("[1,3,4]"),
			want:   "[1,1,2,3,4,4]",
		},
		{
			name:   "test-[1]+[]",
			input1: list.CreateNode("[1]"),
			input2: list.CreateNode("[]"),
			want:   "[1]",
		},
		{
			name:   "test-[]+[1]",
			input1: list.CreateNode("[]"),
			input2: list.CreateNode("[1]"),
			want:   "[1]",
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
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
