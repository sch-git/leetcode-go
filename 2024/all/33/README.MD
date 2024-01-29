# [搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?company_slug=bytedance)

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

可以想到使用 二分的方式：
如果 mid <= left;说明 left-mid 是无序的，mid-right 是有序的；如果 target <= mid，说明 target 落在 left-mid之间
如果 mid > left;说明 left-mid 是有序的，如果 target 满足 nums[left,mid]，则说明 target 在 left-mid中。