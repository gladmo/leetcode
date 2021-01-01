package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/605/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   bool
	}{
		{
			name:   "test-[0]",
			input1: []int{0},
			input2: 1,
			want:   true,
		},
		{
			name:   "test-[1]",
			input1: []int{1},
			input2: 1,
			want:   false,
		},
		{
			name:   "test-[1]",
			input1: []int{1},
			input2: 0,
			want:   true,
		},
		{
			name:   "test-[0]",
			input1: []int{0},
			input2: 0,
			want:   true,
		},
		{
			name:   "test-[]",
			input1: []int{},
			input2: 1,
			want:   false,
		},
		{
			name:   "test-[0,1,0]",
			input1: []int{0, 1, 0},
			input2: 1,
			want:   false,
		},
		{
			name:   "test-[0,0,1,0,1]",
			input1: []int{0, 0, 1, 0, 1},
			input2: 1,
			want:   true,
		},
		{
			name:   "test-[1,0,0,0,1]",
			input1: []int{1, 0, 0, 0, 1},
			input2: 1,
			want:   true,
		},
		{
			name:   "test-[1,0,0,0,1]",
			input1: []int{1, 0, 0, 0, 1},
			input2: 2,
			want:   false,
		},
		{
			name:   "test-[0,1,0,1,0,1,0,0]",
			input1: []int{0, 1, 0, 1, 0, 1, 0, 0},
			input2: 1,
			want:   true,
		},
		{
			name:   "test-[1,0,0,0,0,0,1]",
			input1: []int{1, 0, 0, 0, 0, 0, 1},
			input2: 2,
			want:   true,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input1, test.input2)
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
