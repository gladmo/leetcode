## [搜索二维矩阵 II](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/)

编写一个高效的算法来搜索 `_m_ x _n_` 矩阵 `matrix` 中的一个目标值 `target` 。该矩阵具有以下特性：

*   每行的元素从左到右升序排列。
*   每列的元素从上到下升序排列。

 

**示例 1：**

![image](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid2.jpg)
`
**输入：**matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
**输出：**true
`

**示例 2：**

![image](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/25/searchgrid.jpg)
`
**输入：**matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
**输出：**false
`

 

**提示：**

*   `m == matrix.length`
*   `n == matrix[i].length`
*   `1 
9
 
9
`
*   每行的所有元素从左到右升序排列
*   每列的所有元素从上到下升序排列
*   `-10
9
 
9
`