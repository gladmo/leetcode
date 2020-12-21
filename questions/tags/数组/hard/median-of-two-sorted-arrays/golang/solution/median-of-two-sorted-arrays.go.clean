package solution

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Export(nums1 []int, nums2 []int) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Params: ", nums1, nums2)
			fmt.Println("Panic:", r)
			fmt.Println()
			debug.PrintStack()
			os.Exit(0)
		}
	}()

	return findMedianSortedArrays(nums1, nums2)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
