package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/1797/golang/solution"
)

func main() {
	/*

		"G()(al)"

	*/

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test-G()(al)",
			input: "G()(al)",
			want:  "Goal",
		},
		{
			name:  "test-G()()()()(al)",
			input: "G()()()()(al)",
			want:  "Gooooal",
		},
		{
			name:  "test-(al)G(al)()()G",
			input: "(al)G(al)()()G",
			want:  "alGalooG",
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
