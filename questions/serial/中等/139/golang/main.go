package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/139/golang/solution"
)

func main() {
	/*

			"leetcode"
		["leet","code"]

	*/

	tests := []struct {
		name   string
		input1 string
		input2 []string
		want   bool
	}{
		{
			name:   `test-leetcode-["leet","code"]`,
			input1: "leetcode",
			input2: []string{"leet", "code"},
			want:   true,
		},
		{
			name:   `test-applepenapple-["apple", "pen"]`,
			input1: "applepenapple",
			input2: []string{"apple", "pen"},
			want:   true,
		},
		{
			name:   `test-catsandog-["cats", "dog", "sand", "and", "cat"]`,
			input1: "catsandog",
			input2: []string{"cats", "dog", "sand", "and", "cat"},
			want:   false,
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
