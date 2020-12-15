package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/160/golang/solution"
	"github.com/gladmo/leetcode/utils/list"
)

func main() {

	node1 := list.CreateNode("[4,1,8,4,5]")
	node2 := list.CreateNode("[5,6,1]")
	node2.Next.Next.Next = node1.Next.Next
	tests := []struct {
		name   string
		input1 *solution.ListNode
		input2 *solution.ListNode
		want   bool
	}{
		{
			name:   "test-[4,1,8,4,5]-[5,6,1,8,4,5]",
			input1: node1,
			input2: node2,
			want:   true,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input1)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
