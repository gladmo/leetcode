package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/困难/321/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 []int
		input3 int
		want   []int
	}{
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-1",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 1,
			want:   []int{9},
		},
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-2",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 2,
			want:   []int{9, 8},
		},
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-3",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 3,
			want:   []int{9, 8, 6},
		},
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-4",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 4,
			want:   []int{9, 8, 6, 5},
		},
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-5",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 5,
			want:   []int{9, 8, 6, 5, 3},
		},
		{
			name:   "test-[3,4,6,5]-[9,1,2,5,8,3]-6",
			input1: []int{3, 4, 6, 5},
			input2: []int{9, 1, 2, 5, 8, 3},
			input3: 6,
			want:   []int{9, 8, 4, 6, 5, 3},
		},
		{
			name:   "test-[6,7]-[6,0,4]-5",
			input1: []int{6, 7},
			input2: []int{6, 0, 4},
			input3: 5,
			want:   []int{6, 7, 6, 0, 4},
		},
		{
			name:   "test-[3,9]-[8,9]-3",
			input1: []int{3, 9},
			input2: []int{8, 9},
			input3: 3,
			want:   []int{9, 8, 9},
		},
		{
			name:   "test-[...]-[...]-50",
			input1: []int{8, 0, 4, 4, 1, 7, 3, 6, 5, 9, 3, 6, 6, 0, 2, 5, 1, 7, 7, 7, 8, 7, 1, 4, 4, 5, 4, 8, 7, 6, 2, 2, 9, 4, 7, 5, 6, 2, 2, 8, 4, 6, 0, 4, 7, 8, 9, 1, 7, 0},
			input2: []int{6, 9, 8, 1, 1, 5, 7, 3, 1, 3, 3, 4, 9, 2, 8, 0, 6, 9, 3, 3, 7, 8, 3, 4, 2, 4, 7, 4, 5, 7, 7, 2, 5, 6, 3, 6, 7, 0, 3, 5, 3, 2, 8, 1, 6, 6, 1, 0, 8, 4},
			input3: 50,
			want:   []int{9, 9, 9, 9, 9, 8, 7, 5, 6, 3, 4, 2, 4, 7, 4, 5, 7, 7, 2, 5, 6, 3, 6, 7, 2, 2, 8, 4, 6, 0, 4, 7, 8, 9, 1, 7, 0, 3, 5, 3, 2, 8, 1, 6, 6, 1, 0, 8, 4, 0},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2, test.input3)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2, test.input3)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
