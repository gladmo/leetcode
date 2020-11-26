package tree

import (
	"fmt"
	"strconv"
	"strings"
)

// CreateTree [3,9,20,null,null,15,7]
func CreateTree(treeStr string) (tree *Node) {
	treeStr = strings.TrimSpace(treeStr)
	treeStr = strings.TrimPrefix(treeStr, "[")
	treeStr = strings.TrimSuffix(treeStr, "]")

	item := strings.Split(treeStr, ",")

	// var level = 1
	var rootIndex = 0
	var treeList []*Node
	for i, val := range item {
		val = strings.TrimSpace(val)

		node := CreateNode(val)

		if node != nil {
			treeList = append(treeList, node)
		}

		if i == 0 {
			continue
		}

		currentRoot := treeList[rootIndex]

		if currentRoot == nil {
			continue
		}

		if i%2 == 0 {
			currentRoot.Right = node
			rootIndex++
		} else {
			currentRoot.Left = node
		}
	}

	if len(treeList) > 0 {
		tree = treeList[0]
	}

	return
}

func CreateNode(val string) *Node {
	if val == "null" {
		return nil
	}

	value, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}

	return &Node{
		Val: int(value),
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (th *Node) String() string {
	return fmt.Sprintf("[%s]", strings.Join(LevelOrder(th), ","))
}

func (th *Node) LevelOrder() [][]int {
	var res [][]int
	if th == nil {
		return res
	}

	queue := []*Node{th}

	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})

		var tmp []*Node

		for j := 0; j < len(queue); j++ {
			node := queue[j]

			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		queue = tmp
	}

	return res
}

func LevelOrder(root *Node) []string {
	var res []string
	if root == nil {
		return res
	}

	queue := []*Node{root}

	for i := 0; len(queue) > 0; i++ {

		temp := []*Node{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]

			if node == nil {
				res = append(res, "null")
			} else {
				res = append(res, strconv.Itoa(node.Val))

				if node.Left != nil {
					temp = append(temp, node.Left)
				} else {
					temp = append(temp, nil)
				}

				if node.Right != nil {
					temp = append(temp, node.Right)
				} else {
					temp = append(temp, nil)
				}
			}

		}

		queue = temp
	}

	for res[len(res)-1] == "null" {
		res = res[:len(res)-1]
	}

	return res
}
