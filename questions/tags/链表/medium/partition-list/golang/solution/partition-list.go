package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)


type ListNode struct {
     Val int
     Next *ListNode
}

func Export(head *ListNode, x int) *ListNode {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", head, x)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return partition(head, x)
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
func partition(head *ListNode, x int) *ListNode {

}
