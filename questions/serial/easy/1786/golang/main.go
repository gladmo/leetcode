package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/1786/golang/solution"
)

func main() {
	/*

			"ab"
		["ad","bd","aaab","baa","badab"]

	*/

	tests := []struct {
		name   string
		input1 string
		input2 []string
		want   int
	}{
		{
			name:   `test-ab-["ad","bd","aaab","baa","badab"]`,
			input1: "ab",
			input2: []string{
				"ad", "bd", "aaab", "baa", "badab",
			},
			want: 2,
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
