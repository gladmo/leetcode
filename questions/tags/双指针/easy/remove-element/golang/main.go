package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/27/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   int
	}{
		{
			name:   "test-[1]-1",
			input1: []int{1},
			input2: 1,
			want:   0,
		},
		{
			name:   "test-[2]-3",
			input1: []int{2},
			input2: 3,
			want:   1,
		},
		{
			name:   "test-[2,3]-3",
			input1: []int{2, 3},
			input2: 3,
			want:   1,
		},
		{
			name:   "test-[0,4,4,0,4,4,4,0,2]-4",
			input1: []int{0, 4, 4, 0, 4, 4, 4, 0, 2},
			input2: 4,
			want:   4,
		},
		{
			name:   "test-[3,3]-3",
			input1: []int{3, 3},
			input2: 3,
			want:   0,
		},
		{
			name:   "test-[3,2,2,3]-3",
			input1: []int{3, 2, 2, 3},
			input2: 3,
			want:   2,
		},
		{
			name:   "test-[0,1,2,2,3,0,4,2]-2",
			input1: []int{0, 1, 2, 2, 3, 0, 4, 2},
			input2: 2,
			want:   5,
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
