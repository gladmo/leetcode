package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/136/golang/solution"
)

func main() {
	/*

		[2,2,1]

	*/

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[2,2,1]",
			input: []int{2, 2, 1},
			want:  1,
		},
		{
			name:  "test-[4,1,2,1,2]",
			input: []int{4, 1, 2, 1, 2},
			want:  4,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
