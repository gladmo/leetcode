package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/hard/85/golang/solution"
)

func main() {
	tests := []struct {
		name  string
		input [][]byte
		want  int
	}{
		{
			name: `test-["10100","10111","11111","10010"]`,
			input: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{
			name:  `test-[]`,
			input: [][]byte{},
			want:  0,
		},
		{
			name: `test-["0"]`,
			input: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			name: `test-["1"]`,
			input: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: `test-["00"]`,
			input: [][]byte{
				{'0', '0'},
			},
			want: 0,
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
