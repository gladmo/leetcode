package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/389/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   byte
	}{
		{
			name:   "test-abcd-abcde",
			input1: "abcd",
			input2: "abcde",
			want:   'e',
		},
		{
			name:   "test--y",
			input1: "",
			input2: "y",
			want:   'y',
		},
		{
			name:   "test-a-aa",
			input1: "a",
			input2: "aa",
			want:   'a',
		},
		{
			name:   "test-ae-aea",
			input1: "ae",
			input2: "aea",
			want:   'a',
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
