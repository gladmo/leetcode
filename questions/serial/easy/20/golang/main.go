package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/20/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "test-()",
			input: "()",
			want:  true,
		},
		{
			name:  "test-()[]{}",
			input: "()[]{}",
			want:  true,
		},
		{
			name:  "test-(]",
			input: "(]",
			want:  false,
		},
		{
			name:  "test-([)]",
			input: "([)]",
			want:  false,
		},
		{
			name:  "test-{[]}",
			input: "{[]}",
			want:  true,
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
