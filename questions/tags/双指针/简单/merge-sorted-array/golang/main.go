package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/88/golang/solution"
)

func main() {
	/*

			[1,2,3,0,0,0]
		3
		[2,5,6]
		3

	*/

	tests := []struct {
		name   string
		input1 []int
		input2 int
		input3 []int
		input4 int
		want   []int
	}{
		{
			name:   "test-[1,2,3]-3-[2,5,6]-3",
			input1: []int{1, 2, 3, 0, 0, 0},
			input2: 3,
			input3: []int{2, 5, 6},
			input4: 3,
			want:   []int{1, 2, 2, 3, 5, 6},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			// solution.Export(test.input1, test.input2, test.input3, test.input4)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		solution.Export(test.input1, test.input2, test.input3, test.input4)
		got := test.input1
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
