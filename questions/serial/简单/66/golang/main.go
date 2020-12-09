package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/66/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test-[1,2,3]",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			name:  "test-[9]",
			input: []int{9},
			want:  []int{1, 0},
		},
		{
			name:  "test-[1,9]",
			input: []int{1, 9},
			want:  []int{2, 0},
		},
		{
			name:  "test-[9,9,9,9,9]",
			input: []int{9, 9, 9, 9, 9},
			want:  []int{1, 0, 0, 0, 0, 0},
		},
		{
			name:  "test-[1,9,9]",
			input: []int{1, 9, 9},
			want:  []int{2, 0, 0},
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

		got := solution.Export(test.input)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
