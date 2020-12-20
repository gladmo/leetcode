package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/1812/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  string
	}{

		{
			name:  "test-3-27",
			input: "1-23-45 6",
			want:  "123-456",
		},
		{
			name:  "test-3-27",
			input: "123 4-567",
			want:  "123-45-67",
		},
		{
			name:  "test-3-27",
			input: "123 4-5678",
			want:  "123-456-78",
		},
		{
			name:  "test-3-27",
			input: "12",
			want:  "12",
		},
		{
			name:  "test-3-27",
			input: "--17-5 229 35-39475 ",
			want:  "175-229-353-94-75",
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
