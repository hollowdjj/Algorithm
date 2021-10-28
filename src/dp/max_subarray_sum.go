package dp

import "math"

/*
MaxSubArraySum
最大子序和问题：给定一个整数数组nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回
其最大和。例如：
输入：[-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1]的和最大，为6

这道最大连续子序列和问题和最长递增序列问题非常相似，只不过这里要求的是一个连续子序列。
*/
func MaxSubArraySum(nums []int) int {
	/*
		dp数组的定义为：dp[i]表示以nums[i]结尾的连续子数组的最大和。因此，状态转移方程为：
		        nums[n]                                n = 0
		dp[n] = max{dp[n-1],dp[n-1] + nums[n]}         n > 0
	*/
	dp := make([]int, len(nums))
	//base case
	dp[0] = nums[0]
	//选择
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	//遍历dp得到最大和
	res := math.MinInt
	for _, val := range dp {
		if val > res {
			res = val
		}
	}

	return res
}
