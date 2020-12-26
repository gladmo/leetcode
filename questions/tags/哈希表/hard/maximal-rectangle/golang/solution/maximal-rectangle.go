package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(matrix [][]byte) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", matrix)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return maximalRectangle(matrix)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func maximalRectangle(matrix [][]byte) int {

}
