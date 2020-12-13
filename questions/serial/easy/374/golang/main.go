package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/374/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 int
		input2 int
		want   int
	}{
		{
			name:   "test-10-6",
			input1: 10,
			input2: 6,
			want:   6,
		},
		{
			name:   "test-1-1",
			input1: 1,
			input2: 1,
			want:   1,
		},
		{
			name:   "test-2-1",
			input1: 2,
			input2: 1,
			want:   1,
		},
		{
			name:   "test-2-2",
			input1: 2,
			input2: 2,
			want:   2,
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
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
