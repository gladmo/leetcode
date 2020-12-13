package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/561/golang/solution"
)

func main() {
	/*

		[1,4,3,2]

	*/

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test-[1,4,3,2]",
			input: []int{1, 4, 3, 2},
			want:  4,
		},
		{
			name:  "test-[6,2,6,5,1,2]",
			input: []int{6, 2, 6, 5, 1, 2},
			want:  9,
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
