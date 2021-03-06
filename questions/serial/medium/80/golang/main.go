package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/80/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,1,1,2,2,3]",
			input: []int{1, 1, 1, 2, 2, 3},
			want:  5,
		},
		{
			name:  "test-[0,0,1,1,1,1,2,3,3]",
			input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want:  7,
		},
		{
			name:  "test-[0,0,1,1,1,1,2,3,3,3]",
			input: []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 3},
			want:  7,
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
