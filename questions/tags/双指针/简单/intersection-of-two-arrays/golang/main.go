package main

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/349/golang/solution"
)

func main() {

	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
	}{
		{
			name:   "test-[1,2,2,1]-[2, 2]",
			input1: []int{1, 2, 2, 1},
			input2: []int{2, 2},
			want:   []int{2},
		},
		{
			name:   "test-[4,9,5]-[9,4,9,8,4]",
			input1: []int{4, 9, 5},
			input2: []int{9, 4, 9, 8, 4},
			want:   []int{4, 9},
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
		sort.Ints(got)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
