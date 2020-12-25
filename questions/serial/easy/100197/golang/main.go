package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/100197/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "test-1",
			input: 1,
			want:  1,
		},
		{
			name:  "test-2",
			input: 2,
			want:  2,
		},
		{
			name:  "test-3",
			input: 3,
			want:  4,
		},
		{
			name:  "test-4",
			input: 4,
			want:  7,
		},
		{
			name:  "test-5",
			input: 5,
			want:  13,
		},
		{
			name:  "test-6",
			input: 6,
			want:  24,
		},
		{
			name:  "test-1000000",
			input: 1000000,
			want:  746580045,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input)
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