package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/28/golang/solution"
)

func main() {
	/*

			"hello"
		"ll"

	*/

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   int
	}{
		{
			name:   "test-hello-ll",
			input1: "hello",
			input2: "ll",
			want:   2,
		},
		{
			name:   "test-aaaaa-bba",
			input1: "aaaaa",
			input2: "bba",
			want:   -1,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := test.want
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input1, test.input2)
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
