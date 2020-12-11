package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/14/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  `test-[]`,
			input: []string{},
			want:  "",
		},
		{
			name:  `test-["haha"]`,
			input: []string{"haha"},
			want:  "haha",
		},
		{
			name:  `test-["ab", "a"]`,
			input: []string{"ab", "a"},
			want:  "a",
		},
		{
			name:  `test-["flower","flow","flight"]`,
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			name:  `test-["dog","racecar","car"]`,
			input: []string{"dog", "racecar", "car"},
			want:  "",
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
