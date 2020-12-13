package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/easy/776/golang/solution"
)

func main() {

	tests := []struct {
		name  string
		input *solution.Node
		want  []int
	}{
		{
			name: "test-[1,null,3,2,4,null,5,6]",
			input: &solution.Node{
				Val: 1,
				Children: []*solution.Node{
					{
						Val: 3,
						Children: []*solution.Node{
							{
								Val:      5,
								Children: nil,
							},
							{
								Val:      6,
								Children: nil,
							},
						},
					},
					{
						Val:      2,
						Children: nil,
					},
					{
						Val:      4,
						Children: nil,
					},
				},
			},
			want: []int{5, 6, 3, 2, 4, 1},
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
