package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(flowerbed []int, n int) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", flowerbed, n)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return canPlaceFlowers(flowerbed, n)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func canPlaceFlowers(flowerbed []int, n int) bool {

}
