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

func Export(headA, headB *ListNode) *ListNode {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", headA, headB)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return getIntersectionNode(headA, headB)    
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    
}
