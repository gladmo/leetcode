package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/747/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,100,1,1,1,100,1,1,100,1]",
			input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want:  6,
		},
		{
			name:  "test-[1,100,1,1,1,100,1,1,100]",
			input: []int{1, 100, 1, 1, 1, 100, 1, 1, 100},
			want:  5,
		},
		{
			name:  "test-[1,100,1,1,100,1,1,100]",
			input: []int{1, 100, 1, 1, 100, 1, 1, 100},
			want:  5,
		},
		{
			name:  "test-[10, 15, 20]",
			input: []int{10, 15, 20},
			want:  15,
		},
		{
			name:  "test-[10,15,20,1]",
			input: []int{10, 15, 20, 1},
			want:  16,
		},
		{
			name:  "test-[10, 15]",
			input: []int{10, 15},
			want:  10,
		},
		{
			name:  "test-[15,10]",
			input: []int{15, 10},
			want:  10,
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
