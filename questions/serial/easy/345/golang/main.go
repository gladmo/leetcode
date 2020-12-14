package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/345/golang/solution"
)

func main() {
	/*

		"hello"

	*/

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-hello",
			input: "hello",
			want:  "holle",
		},
		{
			name:  "test-leetcode",
			input: "leetcode",
			want:  "leotcede",
		},
		{
			name:  "test-",
			input: "",
			want:  "",
		},
		{
			name:  "test-aA",
			input: "aA",
			want:  "Aa",
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
