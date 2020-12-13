package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/237/golang/solution"
)

func main() {

	node := &solution.ListNode{Val: 1, Next: &solution.ListNode{Val: 9}}
	head := &solution.ListNode{Val: 4, Next: &solution.ListNode{Val: 5, Next: node}}
	tests := []struct {
		name  string
		input *solution.ListNode
		want  string
	}{
		{
			name:  "test-[4,5,1,9]",
			input: node,
			want:  "[4,5,9]",
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			// solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		solution.Export(test.input)
		got := head.String()
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
