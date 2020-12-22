package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/1000003/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[-2,1,-3,4,-1,2,1,-5,4]",
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want:  6,
		},
		{
			name:  "test-[-2]",
			input: []int{-2},
			want:  -2,
		},
		{
			name:  "test-[]",
			input: []int{},
			want:  0,
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
