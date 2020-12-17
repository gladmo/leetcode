package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/hard/10/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{
			name:   "test-a-ab*a",
			input1: "a",
			input2: "ab*a",
			want:   false,
		},
		{
			name:   "test-aaa-aaaa",
			input1: "aaa",
			input2: "aaaa",
			want:   false,
		},
		{
			name:   "test-aaa-ab*ac*a",
			input1: "aaa",
			input2: "ab*ac*a",
			want:   true,
		},
		{
			name:   "test-ab-.*b",
			input1: "ab",
			input2: ".*b",
			want:   true,
		},
		{
			name:   "test-ab-.*bb",
			input1: "abb",
			input2: ".*bb.*",
			want:   true,
		},
		{
			name:   "test-ab-.*c",
			input1: "ab",
			input2: ".*c",
			want:   false,
		},
		{
			name:   "test-ab-.*ab",
			input1: "ab",
			input2: ".*ab",
			want:   true,
		},
		{
			name:   "test-aab-mis*is*ip*.",
			input1: "mississippi",
			input2: "mis*is*ip*.",
			want:   true,
		},
		{
			name:   "test-aab-c*c*a*b",
			input1: "aab",
			input2: "c*c*a*b",
			want:   true,
		},
		{
			name:   "test-aab-c*******a*b",
			input1: "aab",
			input2: "c*******a*b",
			want:   true,
		},
		{
			name:   "test-aab-c*a*b",
			input1: "aab",
			input2: "c*a*b",
			want:   true,
		},
		{
			name:   "test-aa-aa",
			input1: "aa",
			input2: "aa",
			want:   true,
		},
		{
			name:   "test-aa-aa",
			input1: "aa",
			input2: "a.",
			want:   true,
		},
		{
			name:   "test-aa-aa",
			input1: "aa",
			input2: ".a",
			want:   true,
		},
		{
			name:   "test-aa-aa",
			input1: "aa",
			input2: ".",
			want:   false,
		},
		{
			name:   "test-aa-a",
			input1: "aa",
			input2: "a",
			want:   false,
		},
		{
			name:   "test-aa-a*",
			input1: "aa",
			input2: "a*",
			want:   true,
		},
		{
			name:   "test-ab-.*",
			input1: "ab",
			input2: ".*",
			want:   true,
		},
		{
			name:   "test-mississippi-mis*is*p*.",
			input1: "mississippi",
			input2: "mis*is*p*.",
			want:   false,
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
