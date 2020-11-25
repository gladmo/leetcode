package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/394/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-3[a]2[bc]",
			input: "3[a]2[bc]",
			want:  "aaabcbc",
		},
		{
			name:  "test-3[a2[c]]",
			input: "3[a2[c]]",
			want:  "accaccacc",
		},
		{
			name:  "test-2[abc]3[cd]ef",
			input: "2[abc]3[cd]ef",
			want:  "abcabccdcdcdef",
		},
		{
			name:  "test-abc3[cd]xyz",
			input: "abc3[cd]xyz",
			want:  "abccdcdcdxyz",
		},
		{
			name:  "test-100[leet]",
			input: "100[leet]",
			want:  "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
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
		if got != test.want {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %s, got %s.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
