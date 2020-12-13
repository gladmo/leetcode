package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/50/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 float64
		input2 int
		want   float64
	}{
		{
			name:   "test-2^10",
			input1: 2,
			input2: 10,
			want:   1024,
		},
		{
			name:   "test-2.1^3",
			input1: 2.1,
			input2: 3,
			want:   9.261,
		},
		{
			name:   "test-2^-2",
			input1: 2,
			input2: -2,
			want:   0.25,
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
