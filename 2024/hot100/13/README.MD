# [最大子数组和](https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked)

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。

> 如果f(x-1)<0; 则 f(x) = num[x](num[x]>0)|max(f(x-1),num[x])(num[x]<0)