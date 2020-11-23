package solution


type Node struct {
     Val int
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
    
}
