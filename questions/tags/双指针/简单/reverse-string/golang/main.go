package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/344/golang/solution"
)

func main() {
	/*

		["h","e","l","l","o"]

	*/

	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "test-hello",
			input: []byte("hello"),
			want:  []byte("olleh"),
		},
		{
			name:  "test-HannaH",
			input: []byte("HannaH"),
			want:  []byte("HannaH"),
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			// 执行两次翻转，检测超时时不对结果造成影响
			solution.Export(test.input)
			solution.Export(test.input)

			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		solution.Export(test.input)
		if !reflect.DeepEqual(test.want, test.input) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v).", test.want, test.input))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
