package main

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/困难/212/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 [][]byte
		input2 []string
		want   []string
	}{
		{
			name: "test-[[1],[2],[3],[]]",
			input1: [][]byte{
				[]byte("oaan"),
				[]byte("etae"),
				[]byte("ihkr"),
				[]byte("iflv"),
			},
			input2: []string{
				"oath",
				"pea",
				"eat",
				"rain",
			},
			want: []string{"oath", "eat"},
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
		sort.Strings(got)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
