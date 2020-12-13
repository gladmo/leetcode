package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/454/golang/solution"
)

func main() {
	/*

			[1,2]
		[-2,-1]
		[-1,2]
		[0,2]

	*/

	tests := []struct {
		name                           string
		input1, input2, input3, input4 []int
		want                           int
	}{
		{
			name:   "test-[1,2]-[-2,-1]-[-1,2]-[0,2]",
			input1: []int{1, 2},
			input2: []int{-2, -1},
			input3: []int{-1, 2},
			input4: []int{0, 2},
			want:   2,
		},
		{
			name:   "test-[-1,-1]-[-1,1]-[-1,1]-[1,-1]",
			input1: []int{-1, -1},
			input2: []int{-1, 1},
			input3: []int{-1, 1},
			input4: []int{1, -1},
			want:   6,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2, test.input3, test.input4)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2, test.input3, test.input4)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
