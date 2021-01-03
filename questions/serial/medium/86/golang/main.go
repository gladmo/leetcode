package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/86/golang/solution"
	"github.com/gladmo/leetcode/utils/list"
)

func main() {

	tests := []struct {
		name   string
		input1 *list.Node
		input2 int
		want   string
	}{
		{
			name:   "test-[1,4,3,2,5,2]-3",
			input1: list.CreateNode("[1,4,3,2,5,2]"),
			input2: 3,
			want:   "[1,2,2,4,3,5]",
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input1, test.input2).String()
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
