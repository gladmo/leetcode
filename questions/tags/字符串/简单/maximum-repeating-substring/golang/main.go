package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/1764/golang/solution"
)

func main() {

	tests := []struct {
		name           string
		input1, input2 string
		want           int
	}{
		{
			name:   "test-ababc",
			input1: "ababc",
			input2: "ab",
			want:   2,
		},
		{name: "test-ababc-ba", input1: "ababc", input2: "ba", want: 1},
		{name: "test-ababc-ac", input1: "ababc", input2: "ac", want: 0},
		{name: "test-abababcab-ab", input1: "abababcab", input2: "ab", want: 3},
		{name: "test-a-a", input1: "a", input2: "a", want: 1},
		{name: "test-aaa-a", input1: "aaa", input2: "a", want: 3},
		{name: "test-bbaa-ba", input1: "bbaa", input2: "ba", want: 1},
		{name: "test-babbaabaa-baa", input1: "babbaabaa", input2: "baa", want: 2},
		{name: "test-bbbabbabbab-bba", input1: "bbbabbabbab", input2: "bba", want: 3},
		{name: "test-babbb-bbabb", input1: "babbb", input2: "bbabb", want: 0},
		{name: "test-ababaabba-aaa", input1: "ababaabba", input2: "aaa", want: 0},
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
