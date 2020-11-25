package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/简单/733/golang/solution"
)

func main() {

	tests := []struct {
		name             string
		image            [][]int
		sr, sc, newColor int
		want             [][]int
	}{
		{
			name: "test-[[1,1,1],[1,1,0],[1,0,1]]",
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr: 1, sc: 1, newColor: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "test-[[0,0,0],[0,0,0]]",
			image: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			sr: 0, sc: 0, newColor: 2,
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.image, test.sr, test.sc, test.newColor)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		got := solution.Export(test.image, test.sr, test.sc, test.newColor)
		if !reflect.DeepEqual(test.want, got) {
			testLog.Fail(idx+1, test.name, fmt.Sprintf("want: %v, got %v.", test.want, got))
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
