package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/hard/4/golang/solution"
)

func main() {
	/*

			[1,3]
		[2]

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   float64
	}{
		{
			name:   "test-[1,3]-[2]",
			input1: []int{1, 3},
			input2: []int{2},
			want:   2,
		},
		{
			name:   "test-[1,2]-[3,4]",
			input1: []int{1, 2},
			input2: []int{3, 4},
			want:   2.5,
		},
		{
			name:   "test-[0,0]-[0,0]",
			input1: []int{0, 0},
			input2: []int{0, 0},
			want:   0,
		},
		{
			name:   "test-[]-[1]",
			input1: []int{},
			input2: []int{1},
			want:   1,
		},
		{
			name:   "test-[2]-[]",
			input1: []int{2},
			input2: []int{},
			want:   2,
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
