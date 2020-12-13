package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/334/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "test-[1,2,3,4,5]",
			input: []int{1, 2, 3, 4, 5},
			want:  true,
		},
		{
			name:  "test-[5,4,3,2,1]",
			input: []int{5, 4, 3, 2, 1},
			want:  false,
		},
		{
			name:  "test-[5,1,5,5,2,5,4]",
			input: []int{5, 1, 5, 5, 2, 5, 4},
			want:  true,
		},
		{
			name:  "test-[5,1,5,4,5,2,5,4]",
			input: []int{5, 1, 5, 4, 5, 2, 5, 4},
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
