package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/217/golang/solution"
)

func main() {
	/*

		[1,2,3,1]

	*/

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "test-[1,2,3,1]",
			input: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			name:  "test-[1,2,3,4]",
			input: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			name:  "test-[1,1,1,3,3,4,3,2,4,2]",
			input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want:  true,
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
