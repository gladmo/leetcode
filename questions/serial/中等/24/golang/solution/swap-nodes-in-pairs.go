package solution

import (
	"github.com/gladmo/leetcode/utils/list"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
type ListNode = list.Node

func Export(head *ListNode) *ListNode {
	return swapPairs(head)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

}
