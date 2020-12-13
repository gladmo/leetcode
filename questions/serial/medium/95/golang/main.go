package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/95/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input int
		want  []string
	}{
		{
			name:  "test-3",
			input: 3,
			want: []string{
				"[1,null,2,null,3]",
				"[1,null,3,2]",
				"[2,1,3]",
				"[3,1,null,null,2]",
				"[3,2,null,1]",
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
		for idx, node := range got {
			if !reflect.DeepEqual(test.want[idx], node.String()) {
				testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want[idx], node.String()))
				continue
			}
		}

		testLog.Pass(idx+1, test.name)
	}
}
