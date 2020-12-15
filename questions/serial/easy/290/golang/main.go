package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/290/golang/solution"
)

func main() {
	/*

			"abba"
		"dog cat cat dog"

	*/

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{
			name:   "test-abba-dog cat cat dog",
			input1: "abba",
			input2: "dog cat cat dog",
			want:   true,
		},
		{
			name:   "test-abba-dog cat cat fish",
			input1: "abba",
			input2: "dog cat cat fish",
			want:   false,
		},
		{
			name:   "test-aaaa-dog cat cat dog",
			input1: "aaaa",
			input2: "dog cat cat dog",
			want:   false,
		},
		{
			name:   "test-abba-dog dog dog dog",
			input1: "abba",
			input2: "dog dog dog dog",
			want:   false,
		},
		{
			name:   "test-aaaa-dog dog dog dog",
			input1: "aaaa",
			input2: "dog dog dog dog",
			want:   true,
		},
		{
			name:   "test-aba-cat cat cat dog",
			input1: "aba",
			input2: "cat cat cat dog",
			want:   false,
		},
		{
			name:   "test-abc-a b c",
			input1: "abc",
			input2: "b c a",
			want:   true,
		},
		{
			name:   "test-abab-dog cat cat dog",
			input1: "abab",
			input2: "dog cat cat dog",
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
