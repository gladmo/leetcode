## [单词拆分](https://leetcode-cn.com/problems/word-break/)

给定一个**非空**字符串 _s_ 和一个包含**非空**单词的列表 _wordDict_，判定 _s_ 是否可以被空格拆分为一个或多个在字典中出现的单词。

**说明：**

*   拆分时可以重复使用字典中的单词。
*   你可以假设字典中没有重复的单词。

**示例 1：**

`**输入:** s = "leetcode", wordDict = ["leet", "code"]
**输出:** true
**解释:** 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
`

**示例 2：**

`**输入:** s = "applepenapple", wordDict = ["apple", "pen"]
**输出:** true
**解释:** 返回 true 因为 `"`applepenapple`"` 可以被拆分成 `"`apple pen apple`"`。
     注意你可以重复使用字典中的单词。
`

**示例 3：**

`**输入:** s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
**输出:** false
`