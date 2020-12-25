package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(g []int, s []int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", g, s)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return findContentChildren(g, s)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func findContentChildren(g []int, s []int) int {

}
