package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/70/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "test-1-1",
			input: 1,
			want:  1,
		},
		{
			name:  "test-2-2",
			input: 2,
			want:  2,
		},
		{
			name:  "test-3-3",
			input: 3,
			want:  3,
		},
		{
			name:  "test-4-5",
			input: 4,
			want:  5,
		},
		{
			name:  "test-5-8",
			input: 5,
			want:  8,
		},
		{
			name:  "test-44-1134903170",
			input: 44,
			want:  1134903170,
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
