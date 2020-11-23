package main

import (
	"fmt"

	"github.com/gladmo/leetcode/questions/serial/中等/133/golang/solution"
)

func main() {
	/*

		[[2,4],[1,3],[2,4],[1,3]]

	*/

	node1 := &solution.Node{
		Val:       1,
		Neighbors: nil,
	}
	node2 := &solution.Node{
		Val:       2,
		Neighbors: nil,
	}

	node1.Neighbors = []*solution.Node{node2}
	node2.Neighbors = []*solution.Node{node1}

	tests := []struct {
		name string
		node *solution.Node
	}{
		{
			name: "test-nil",
			node: &solution.Node{},
		},
		{
			name: "test-[[2],[1]]",
			node: node1,
		},
	}

	for idx, test := range tests {
		if solution.Export(test.node) == test.node {
			Fail(idx+1, len(tests), test.name, "shallow copy")
			continue
		}

		Pass(idx+1, len(tests), test.name)
	}
}

func Fail(index, length int, name, reason string) {
	fmt.Println(
		fmt.Sprintf(
			"[%d/%d] Test Name: %s, Failed, Reason: %s",
			index,
			length,
			name,
			reason,
		))
}

func Pass(index, length int, name string) {
	fmt.Println(
		fmt.Sprintf(
			"[%d/%d] name: %s, Pass",
			index,
			length,
			name,
		))
}
