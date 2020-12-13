package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/782/golang/solution"
)

func main() {
	/*

			"aA"
		"aAAbbbb"

	*/

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   int
	}{
		{
			name:   "test-aA-aAAbbbb",
			input1: "aA",
			input2: "aAAbbbb",
			want:   3,
		},
		{
			name:   "test-z-ZZ",
			input1: "z",
			input2: "ZZ",
			want:   0,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.input1, test.input2)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.input1, test.input2)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
