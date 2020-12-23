package str

// Sequence 用于查询该字符串是否存在目录子串
// 注意: 只实现全小写字母匹配
//
// 原理:
// 创建一个字符串长度二维数组，倒序遍历字符串，当前字符出现
// 则当前位置设置为字符串索引，其它位置设置为下一层的值
//
// 使用:
// seq := New("ahbgdc")
//
// true == seq.HasSubsequence("abc")
type Sequence struct {
	length int
	dp     [][]int
}

// NewSequence 创建匹配对象
func NewSequence(str string) *Sequence {
	length := len(str)
	var dp = make([][]int, length+1)

	lastRow := make([]int, 26)
	for i := 0; i < 26; i++ {
		lastRow[i] = length
	}

	dp[length] = lastRow

	for i := length - 1; i >= 0; i-- {
		row := make([]int, 26)

		for j := 0; j < 26; j++ {
			if j == int(str[i]-'a') {
				row[j] = i
			} else {
				row[j] = dp[i+1][j]
			}
		}

		dp[i] = row
	}

	return &Sequence{dp: dp, length: length}
}

// HasSubsequence 是否包含子串 target
func (th *Sequence) HasSubsequence(target string) bool {
	var idx int
	for i := 0; i < len(target); i++ {
		val := th.dp[idx][int(target[i]-'a')]
		if val == th.length {
			return false
		}
		idx = val + 1
	}
	return true
}
