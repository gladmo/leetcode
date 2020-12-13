## [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)

给你两个有序整数数组 _nums1_ 和 _nums2_，请你将 _nums2_ 合并到 _nums1_中_，_使 _nums1_ 成为一个有序数组。

 

**说明：**

*   初始化 _nums1_ 和 _nums2_ 的元素数量分别为 _m_ 和 _n_ 。
*   你可以假设 _nums1_有足够的空间（空间大小大于或等于 _m + n_）来保存 _nums2_ 中的元素。

 

**示例：**

`
**输入：**
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

**输出：**[1,2,2,3,5,6]`

 

**提示：**

*   `-10^9 <= nums1[i], nums2[i] <= 10^9`
*   `nums1.length == m + n`
*   `nums2.length == n`