package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/153/golang/solution"
)

func main() {
	/*

		[3,4,5,1,2]

	*/

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[3,4,5,1,2]",
			input: []int{3, 4, 5, 1, 2},
			want:  1,
		},
		{
			name:  "test-[4,5,6,7,0,1,2]",
			input: []int{4, 5, 6, 7, 0, 1, 2},
			want:  0,
		},
		{
			name:  "test-[4,5,6,7,0]",
			input: []int{4, 5, 6, 7, 0},
			want:  0,
		},
		{
			name:  "test-[4,5,6,7,0,1]",
			input: []int{4, 5, 6, 7, 0, 1},
			want:  0,
		},
		{
			name:  "test-[11,13,15,17]",
			input: []int{11, 13, 15, 17},
			want:  11,
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
