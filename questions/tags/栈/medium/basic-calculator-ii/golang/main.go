package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/227/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test-1*2-3/4+5*6-7*8+9/10",
			input: "1*2-3/4+5*6-7*8+9/10",
			want:  -24,
		},
		{
			name:  "test-3+2*2",
			input: "3+2*2",
			want:  7,
		},
		{
			name:  "test-3/2",
			input: "3/2",
			want:  1,
		},
		{
			name:  "test-3/2",
			input: " 3/2 ",
			want:  1,
		},
		{
			name:  "test- 3+5 / 2 ",
			input: " 3+5 / 2 ",
			want:  5,
		},
		{
			name:  "test- ... ",
			input: "3+5 / 2 + 8   + 2 * 9/3 + 3 * 5 + 2 + 5 -12 + 3 * 5 * 5 +2 * 5 +10 -23 -3 -5 + 10 / 3",
			want:  96,
		},
		{
			name:  "test- ... ",
			input: "5 -12 +2 * 5 +10 -2",
			want:  11,
		},
		{
			name:  "test-4/3+2",
			input: "4/3+2",
			want:  3,
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
