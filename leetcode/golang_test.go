package leetcode

import (
	"fmt"
	"testing"
)

func Test_parseGoCode(t *testing.T) {
	code := `
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

}
`

	fmt.Println(parseGoCode(code))
}
