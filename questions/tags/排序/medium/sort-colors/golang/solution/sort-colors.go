package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(nums []int)  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", nums)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	sortColors(nums)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func sortColors(nums []int)  {

}
