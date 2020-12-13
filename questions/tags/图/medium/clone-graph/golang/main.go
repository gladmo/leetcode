package main

import (
	"context"
	"time"

	"github.com/gladmo/leetcode/leet"
	"github.com/gladmo/leetcode/questions/serial/medium/133/golang/solution"
)

func main() {
	// [[2,4],[1,3],[2,4],[1,3]]
	node1234 := &solution.Node{Val: 1}
	node2 := &solution.Node{Val: 2}
	node3 := &solution.Node{Val: 3}
	node4 := &solution.Node{Val: 4}
	node1234.Neighbors = []*solution.Node{node2, node4}
	node2.Neighbors = []*solution.Node{node1234, node3}
	node3.Neighbors = []*solution.Node{node2, node4}
	node4.Neighbors = []*solution.Node{node1234, node3}

	// [[2],[1]]
	node11 := &solution.Node{Val: 1}
	node12 := &solution.Node{Val: 2}
	node11.Neighbors = []*solution.Node{node12}
	node12.Neighbors = []*solution.Node{node11}

	tests := []struct {
		name string
		node *solution.Node
	}{
		{
			name: "test-nil",
			node: nil,
		},
		{
			name: "test-empty",
			node: &solution.Node{},
		},
		{
			name: "test-[[2],[1]]",
			node: node11,
		},
		{
			name: "test-[[2,4],[1,3],[2,4],[1,3]]",
			node: node1234,
		},
	}

	testLog := leet.NewTestLog(len(tests))
	defer testLog.Render()

	timeoutDuration := time.Second * 2

	for idx, test := range tests {
		// 超时检测
		timeout := leet.Timeout(timeoutDuration, func(ctx context.Context, cancel context.CancelFunc) {
			solution.Export(test.node)
			cancel()
		})

		if timeout {
			testLog.Fail(idx+1, test.name, "timeout")
			continue
		}

		if solution.Export(test.node) == test.node && test.node != nil {
			testLog.Fail(idx+1, test.name, "shallow copy")
			continue
		}

		testLog.Pass(idx+1, test.name)
	}
}
