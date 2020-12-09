package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/283/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test-[0]",
			input: []int{0},
			want:  []int{0},
		},
		{
			name:  "test-[1]",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "test-[0,1]",
			input: []int{0, 1},
			want:  []int{1, 0},
		},
		{
			name:  "test-[0,1,0,3,12]",
			input: []int{0, 1, 0, 3, 12},
			want:  []int{1, 3, 12, 0, 0},
		},
		{
			name:  "test-[0,0,1,0,3,12,0,1,0,0,0]",
			input: []int{0, 0, 1, 0, 3, 12, 0, 1, 0, 0, 0},
			want:  []int{1, 3, 12, 1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:  "test-[1,2]",
			input: []int{1, 2},
			want:  []int{1, 2},
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

		solution.Export(test.input)
		got := test.input
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
