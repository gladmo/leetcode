package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/367/golang/solution"
)

func main() {
	/*

		16

	*/

	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "test-16",
			input: 16,
			want:  true,
		},
		{
			name:  "test-9",
			input: 9,
			want:  true,
		},
		{
			name:  "test-14",
			input: 14,
			want:  false,
		},
		{
			name:  "test-15",
			input: 15,
			want:  false,
		},
		{
			name:  "test-225",
			input: 225,
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
