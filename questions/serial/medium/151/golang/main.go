package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/151/golang/solution"
)

func main() {
	/*

		"the sky is blue"

	*/

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-the sky is blue",
			input: "the sky is blue",
			want:  "blue is sky the",
		},
		{
			name:  "test-  hello world!  ",
			input: "  hello world!  ",
			want:  "world! hello",
		},
		{
			name:  "test-a good   example",
			input: "a good   example",
			want:  "example good a",
		},
		{
			name:  "test-  Bob    Loves  Alice   ",
			input: "  Bob    Loves  Alice   ",
			want:  "Alice Loves Bob",
		},
		{
			name:  "test-Alice does not even like bob",
			input: "Alice does not even like bob",
			want:  "bob like even not does Alice",
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
