package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/392/golang/solution"
)

func main() {
	/*

			"abc"
		"ahbgdc"

	*/

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{
			name:   "test-abc-ahbgdc",
			input1: "abc",
			input2: "ahbgdc",
			want:   true,
		},
		{
			name:   "test-axc-ahbgdc",
			input1: "axc",
			input2: "ahbgdc",
			want:   false,
		},
		{
			name:   "test--ahbgdc",
			input1: "",
			input2: "ahbgdc",
			want:   true,
		},
		{
			name:   "test--",
			input1: "",
			input2: "",
			want:   true,
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
