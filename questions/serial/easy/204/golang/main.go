package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/204/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "test-10",
			input: 10,
			want:  4,
		},
		{
			name:  "test-0",
			input: 0,
			want:  0,
		},
		{
			name:  "test-1",
			input: 1,
			want:  0,
		},
		{
			name:  "test-100",
			input: 100,
			want:  25,
		},
		{
			name:  "test-2",
			input: 2,
			want:  0,
		},
		{
			name:  "test-499979",
			input: 499979,
			want:  41537,
		},
		{
			name:  "test-1500000",
			input: 1500000,
			want:  114155,
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
