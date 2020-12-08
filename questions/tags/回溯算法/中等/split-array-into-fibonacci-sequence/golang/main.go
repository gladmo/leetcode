package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/中等/872/golang/solution"
)

func main() {
	/*

		"123456579"

	*/

	tests := []struct {
		name  string
		input string
		want  [][]int
	}{
		{
			name:  "test-123456579",
			input: "123456579",
			want: [][]int{
				{123, 456, 579},
			},
		},
		{
			name:  "test-11235813",
			input: "11235813",
			want: [][]int{
				{1, 1, 2, 3, 5, 8, 13},
			},
		},
		{
			name:  "test-112358130",
			input: "112358130",
			want:  [][]int{},
		},
		{
			name:  "test-0123",
			input: "0123",
			want:  [][]int{},
		},
		{
			name:  "test-1101111",
			input: "1101111",
			want: [][]int{
				{110, 1, 111},
				{11, 0, 11, 11},
			},
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
		var find bool
		for _, want := range test.want {
			if reflect.DeepEqual(want, got) {
				find = true
			}
		}
		if len(test.want) != 0 && !find {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
