package solution

import (
	"github.com/gladmo/leetcode/utils/tree"
)

// type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
// }

type TreeNode = tree.Node

func Export(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {

}
