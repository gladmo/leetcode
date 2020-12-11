package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(s string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", s)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return longestPalindrome(s)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func longestPalindrome(s string) string {

}
