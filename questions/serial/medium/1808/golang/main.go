package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/1808/golang/solution"
)

func main() {
	/*

		[5,3,1,4,2]

	*/

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[5,3,1,4,2]",
			input: []int{5, 3, 1, 4, 2},
			want:  6,
		},
		{
			name:  "test-[7,90,5,1,100,10,10,2]",
			input: []int{7, 90, 5, 1, 100, 10, 10, 2},
			want:  122,
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
