package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/5/golang/solution"
)

func main() {
	/*

		"babad"

	*/

	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "test-babad",
			input: "babad",
			want:  []string{"aba", "bab"},
		},
		{
			name:  "test-cbbd",
			input: "cbbd",
			want:  []string{"bb"},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		got := strings.Join(test.want, ",")
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			got = solution.Export(test.input)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		if !strings.Contains(strings.Join(test.want, ","), got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
