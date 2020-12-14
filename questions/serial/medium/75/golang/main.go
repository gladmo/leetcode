package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/75/golang/solution"
)

func main() {
	/*

		[2,0,2,1,1,0]

	*/

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test-[2,0,2,1,1,0]",
			input: []int{2, 0, 2, 1, 1, 0},
			want:  []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:  "test-[2,0,1]",
			input: []int{2, 0, 1},
			want:  []int{0, 1, 2},
		},
		{
			name:  "test-[0]",
			input: []int{0},
			want:  []int{0},
		},
		{
			name:  "test-[1]",
			input: []int{1},
			want:  []int{1},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.input
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input)
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
