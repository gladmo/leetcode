package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/387/golang/solution"
)

func main() {
	/*

		"leetcode"

	*/

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "test-leetcode",
			input: "leetcode",
			want:  0,
		},
		{
			name:  "test-loveleetcode",
			input: "loveleetcode",
			want:  2,
		},
		{
			name:  "test-z",
			input: "z",
			want:  0,
		},
		{
			name:  "test-",
			input: "",
			want:  -1,
		},
		{
			name:  "test-bd",
			input: "bd",
			want:  0,
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
