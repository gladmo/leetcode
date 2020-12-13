package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/179/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{
			name:  "test-[3,3,33,333,33333,333,33]",
			input: []int{3, 3, 33, 333, 33333, 343, 33},
			want:  "34333333333333333",
		},
		{
			name:  "test-[111311,1113]",
			input: []int{111311, 1113},
			want:  "1113111311",
		},
		{
			name:  "test-[10,2]",
			input: []int{10, 2},
			want:  "210",
		},
		{
			name:  "test-[10,3,2]",
			input: []int{10, 3, 2},
			want:  "3210",
		},
		{
			name:  "test-[3,30,34,5,9]",
			input: []int{3, 30, 34, 5, 9},
			want:  "9534330",
		},
		{
			name:  "test-[0]",
			input: []int{0},
			want:  "0",
		},
		{
			name:  "test-[0,0]",
			input: []int{0, 0},
			want:  "0",
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			// solution.Export(test.input)
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
