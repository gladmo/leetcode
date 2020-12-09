package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/890/golang/solution"
)

func main() {
	/*

		[5,5,5,10,20]

	*/

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "test-[5,5,5,10,20]",
			input: []int{5, 5, 5, 10, 20},
			want:  true,
		},
		{
			name:  "test-[5,5,10]",
			input: []int{5, 5, 10},
			want:  true,
		},
		{
			name:  "test-[10,10]",
			input: []int{10, 10},
			want:  false,
		},
		{
			name:  "test-[5,5,10,10,20]",
			input: []int{5, 5, 10, 10, 20},
			want:  false,
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
