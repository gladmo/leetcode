## [存在重复元素 III](https://leetcode-cn.com/problems/contains-duplicate-iii/)

在整数数组 `nums` 中，是否存在两个下标 **_i_** 和 **_j_**，使得 **nums [i]** 和 **nums [j]** 的差的绝对值小于等于 _**t**_ ，且满足 **_i_** 和 **_j_** 的差的绝对值也小于等于 _**ķ**_ 。

如果存在则返回 `true`，不存在返回 `false`。

 

**示例 1:**

`**输入:** nums = [1,2,3,1], k __ = 3, t = 0
**输出:** true`

**示例 2:**

`**输入:** nums = [1,0,1,1], k __ = __ 1, t = 2
**输出:** true`

**示例 3:**

`**输入:** nums = [1,5,9,1,5,9], k = 2, t = 3
**输出:** false`