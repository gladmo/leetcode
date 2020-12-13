package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/205/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 string
		input2 string
		want   bool
	}{
		{
			name:   "test-egg-add",
			input1: "egg",
			input2: "add",
			want:   true,
		},
		{
			name:   "test-foo-bar",
			input1: "foo",
			input2: "bar",
			want:   false,
		},
		{
			name:   "test-paper-title",
			input1: "paper",
			input2: "title",
			want:   true,
		},
		{
			name:   "test-ab-aa",
			input1: "ab",
			input2: "aa",
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
