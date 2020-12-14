package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(N int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", N)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return monotoneIncreasingDigits(N)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func monotoneIncreasingDigits(N int) int {

}
