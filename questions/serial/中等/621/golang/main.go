package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/621/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []byte
		input2 int
		want   int
	}{
		{
			name:   `test-AAABBB`,
			input1: []byte("AAABBB"),
			input2: 2,
			want:   8,
		},
		{
			name:   `test-AAABBB`,
			input1: []byte("AAABBB"),
			input2: 50,
			want:   104,
		},
		{
			name:   `test-AAABBBCCCDDE`,
			input1: []byte("AAABBBCCCDDE"),
			input2: 2,
			want:   12,
		},
		{
			name:   `test-AAAAAABCDEFG`,
			input1: []byte("AAAAAABCDEFG"),
			input2: 2,
			want:   16,
		},
		{
			name:   `test-AAABBB`,
			input1: []byte("AAABBB"),
			input2: 0,
			want:   6,
		},
		{
			name:   `test-AAABBBCC`,
			input1: []byte("AAABBBCC"),
			input2: 0,
			want:   8,
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
