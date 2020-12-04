package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/451/golang/solution"
	"github.com/gladmo/leetcode/utils/str"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-tree",
			input: "tree",
			want:  "eetr",
		},
		{
			name:  "test-cccaaa",
			input: "cccaaa",
			want:  "cccaaa",
		},
		{
			name:  "test-Aabb",
			input: "Aabb",
			want:  "bbAa",
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
		got = str.Sort(got)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
