package list

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

// CreateNode [1,2,3,4]
func CreateNode(treeStr string) *Node {
	treeStr = strings.TrimSpace(treeStr)
	treeStr = strings.TrimPrefix(treeStr, "[")
	treeStr = strings.TrimSuffix(treeStr, "]")

	var root *Node
	var current *Node

	item := strings.Split(treeStr, ",")
	for idx, val := range item {
		if val == "" {
			continue
		}

		value, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			panic(err)
		}

		newNode := &Node{Val: int(value)}

		if idx == 0 {
			root = newNode
			current = newNode
			continue
		}

		current.Next = newNode
		current = newNode
	}

	return root
}

func (th *Node) String() string {
	return fmt.Sprintf("[%s]", strings.Join(th.Values(), ","))
}

func (th *Node) Values() []string {
	var res []string
	if th == nil {
		return res
	}

	res = append(res, fmt.Sprint(th.Val))
	if th.Next != nil {
		res = append(res, th.Next.Values()...)
	}

	return res
}
