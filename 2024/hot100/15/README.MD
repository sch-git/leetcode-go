# [轮转数组](https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-100-liked)

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

> 如果k == len(nums)，那么数组没有变化；所以先判断 k = k%len(nums) != 0
> 从后向前依次复制 num[len(num)-1] = num[len(-1)-k];将最后k个位置的数保存，赋值到数组最前