package solution

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/gladmo/leetcode/utils/tree"
)

// type TreeNode struct {
//      Val int
//      Left *TreeNode
//      Right *TreeNode
// }

type TreeNode = tree.Node

func Export(root *TreeNode) [][]int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", root)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return zigzagLevelOrder(root)
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
func zigzagLevelOrder(root *TreeNode) [][]int {

}
