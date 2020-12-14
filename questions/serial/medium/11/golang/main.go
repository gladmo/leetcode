package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/11/golang/solution"
)

func main() {
	/*

		[1,8,6,2,5,4,8,3,7]

	*/

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,8,6,2,5,4,8,3,7]",
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
		{
			name:  "test-[1,1]",
			input: []int{1, 1},
			want:  1,
		},
		{
			name:  "test-[4,3,2,1,4]",
			input: []int{4, 3, 2, 1, 4},
			want:  16,
		},
		{
			name:  "test-[1,2,1]",
			input: []int{1, 2, 1},
			want:  2,
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
