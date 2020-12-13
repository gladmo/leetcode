package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/167/golang/solution"
)

func main() {
	/*

			[2,7,11,15]
		9

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "test-[2,7,11,15]-9",
			input1: []int{2, 7, 11, 15},
			input2: 9,
			want:   []int{1, 2},
		},
		{
			name:   "test-[2,7,11,15]-17",
			input1: []int{2, 7, 11, 15},
			input2: 17,
			want:   []int{1, 4},
		},
		{
			name:   "test-[2,7,11,15]-22",
			input1: []int{2, 7, 11, 15},
			input2: 22,
			want:   []int{2, 4},
		},
		{
			name:   "test-[2,7,11,15]-18",
			input1: []int{2, 7, 11, 15},
			input2: 18,
			want:   []int{2, 3},
		},
		{
			name:   "test-[2,7,11,15]-26",
			input1: []int{2, 7, 11, 15},
			input2: 26,
			want:   []int{3, 4},
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
