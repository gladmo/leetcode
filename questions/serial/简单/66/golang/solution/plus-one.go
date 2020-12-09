package solution

func Export(digits []int) []int {
	return plusOne(digits)
}

/****************************************************/
/******** 以下为 Leetcode 示例部分（提交PR请还原） *******/
/******** 使用 (./leetcode clear) 初始化所有问题 *******/
/****************************************************/

func plusOne(digits []int) []int {
	i := len(digits) - 1

	digits[i]++
	for digits[i] > 9 {
		digits[i] = 0
		if i-1 < 0 {
			digits = append([]int{1}, digits...)
		} else {
			i--
			digits[i]++
		}
	}

	return digits
}
