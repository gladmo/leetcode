package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/1000023/golang/solution"
)

func main() {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,3,1]",
			input: []int{1, 3, 1},
			want:  3,
		},
		{
			name:  "test-[1,2,3,1]",
			input: []int{1, 2, 3, 1},
			want:  4,
		},
		{
			name:  "test-[2,1,4,5,3,1,1,3]",
			input: []int{2, 1, 4, 5, 3, 1, 1, 3},
			want:  12,
		},
		{
			name:  "test-[2,7,9,3,1]",
			input: []int{2, 7, 9, 3, 1},
			want:  12,
		},
		{
			name:  "test-[2]",
			input: []int{2},
			want:  2,
		},
		{
			name:  "test-[2,7]",
			input: []int{2, 7},
			want:  7,
		},
		{
			name:  "test-[]",
			input: []int{},
			want:  0,
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
