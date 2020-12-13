package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/34/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "test-[5,7,7,8,8,10]-8",
			input1: []int{5, 7, 7, 8, 8, 10},
			input2: 8,
			want:   []int{3, 4},
		},
		{
			name:   "test-[5,7,7,8,8,10]-6",
			input1: []int{5, 7, 7, 8, 8, 10},
			input2: 6,
			want:   []int{-1, -1},
		},
		{
			name:   "test-[]-0",
			input1: []int{},
			input2: 0,
			want:   []int{-1, -1},
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
