package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/376/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,7,4,9,2,5]",
			input: []int{1, 7, 4, 9, 2, 5},
			want:  6,
		},
		{
			name:  "test-[1,7]",
			input: []int{1, 7},
			want:  2,
		},
		{
			name:  "test-[1]",
			input: []int{1},
			want:  1,
		},
		{
			name:  "test-[]",
			input: []int{},
			want:  0,
		},
		{
			name:  "test-[1,17,5,10,13,15,10,5,16,8]",
			input: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			want:  7,
		},
		{
			name:  "test-[1,2,3,4,5,6,7,8,9]",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:  2,
		},
		{
			name:  "test-[5,5,5,5,5,5]",
			input: []int{5, 5, 5, 5, 5, 5},
			want:  1,
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
