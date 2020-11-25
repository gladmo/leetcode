package solution

type Node struct {
	Val       int
	Neighbors []*Node
}

func Export(node *Node) *Node {
	return cloneGraph(node)
}

/**************************************/
/******** 以下为 Leetcode 源码部分 *******/
/**************************************/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	var scanned = make(map[int]*Node)
	var newNodes = make(map[int]*Node)

	queue := []*Node{node}

	for i := 0; len(queue) > 0; i++ {
		var tmp []*Node

		for j := 0; j < len(queue); j++ {
			current := queue[j]

			_, ok := scanned[current.Val]
			if !ok {
				tmp = append(tmp, current.Neighbors...)
				scanned[current.Val] = current
				newNodes[current.Val] = &Node{Val: current.Val}
			}
		}

		queue = tmp
	}

	for _, n := range scanned {
		newNode := newNodes[n.Val]
		for _, neighbor := range n.Neighbors {
			neighborNode := newNodes[neighbor.Val]

			newNode.Neighbors = append(newNode.Neighbors, neighborNode)
		}
	}

	return newNodes[node.Val]
}
