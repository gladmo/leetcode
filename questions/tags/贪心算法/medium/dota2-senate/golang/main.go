package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/649/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-RD",
			input: "RD",
			want:  "Radiant",
		},
		{
			name:  "test-RDD",
			input: "RDD",
			want:  "Dire",
		},
		{
			name:  "test-RDDD",
			input: "RDDD",
			want:  "Dire",
		},
		{
			name:  "test-RDDR",
			input: "RDDR",
			want:  "Radiant",
		},
		{
			name:  "test-RRDDD",
			input: "RRDDD",
			want:  "Radiant",
		},
		{
			name:  "test-DRRDRDRDRDDRDRDR",
			input: "DRRDRDRDRDDRDRDR",
			want:  "Radiant",
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
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
