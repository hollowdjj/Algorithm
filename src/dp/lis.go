package dp

import "math"

/*
最长递增子序列问题(Longest Increasing Subsequence)：给定一个无序数组，找到其中的最长递增子序列。
例如数组[10,9,2,5,3,7,101,18]的最长递增子序列为[2,3,7,101]长度为4。

动态规划的核心思想是数学归纳法，也就是说，对一个dp数组，我们可以假设在dp[0,1,...n-1]都是已知的情况下
如何计算得到dp[n]。以最长递增子序列问题为例，dp数组可以这样定义：dp[i]表示以num[i]结尾的最长递增子序
列的长度。如此定义后，就有dp[i] = max{dp[j],其中0<=j<i且num[j] < num[i]} + 1，也即状态转移方程。
显然，这里动态规划求解的时间复杂度为O(n^2)(2层循环嵌套)

从这道题可以看出，动态规划问题的关键在于dp数组的定义。随后，需要思考如何在已知dp[0...n-1]的情况下计算
出dp[n](即选择)
*/

func Lis(nums []int) int {
	dp := make([]int, len(nums))
	//base case dp[i]的初始值均为1，因为以num[i]结尾的最长递增子序列一定包含num[i]
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	//选择
	res := math.MinInt
	for i := 0; i < len(dp); i++ {
		//找到dp[0...i-1]中结尾比nums[i]小的最长序列的长度
		length := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && length < dp[j] {
				length = dp[j]
			}
		}
		dp[i] = length + 1
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
