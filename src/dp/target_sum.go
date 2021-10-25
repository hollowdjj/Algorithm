package dp

/*
目标和问题：给定一个非负整数数组a1,a2...an，以及一个目标数S。现在你有两个符号+和-。对于数组中的
任意一个整数，你都可以从+或-中选择一个符号添加在前面。编写一个函数，返回可以使最终数组和为目标数
S的所有添加符号的方法数量。例如：
num = [1,1,1,1,1] S = 3	共有5中方法
*/

/*
首先使用回溯法求解这个目标和问题。回缩法的模板如下：
func BackTrack(路径，层数) {
	if 满足结束条件 {
		添加结果或更新结果
		return
	}
	for 选择 in 选择列表 {
		做选择
		BackTrack(路径，选择列表)    //沿着选择的路径一直往下走
		撤销选择
    }
}
时间复杂度为O(2^n)，因为这里没有剪枝的过程，所以其实就是暴力枚举。
*/
var result = 0

func TargetSumWaysBackTrack(nums []int, s int) int {
	if len(nums) == 0 {
		return 0
	}

	backTrack(nums, 0, s)
	return result
}

func backTrack(nums []int, i int, rest int) {
	//递归结束条件
	if i == len(nums) {
		if rest == 0 {
			result += 1
		}
		return
	}

	//首先选择在nums[i]前面添加“+”
	rest -= nums[i]
	backTrack(nums, i+1, rest)
	//剪枝
	if nums[i] == 0 {
		return
	}
	//撤销选择
	rest += nums[i]
	//然后选择在nums[i]前面添加“-”
	rest += nums[i]
	backTrack(nums, i+1, rest)
}

/*
此题可用动态规划求解，因为其可转换成一个子集划分问题。首先，把nums中的数划分成A和B两个子集。
其中A集合中全部为添加“+”的数，而B集合中全部为添加“-”的数。从而有：
sum(A) - sum(B) = target
sum(A) + sum(B) = sum(nums)
sum(A) = (target + sum(nums)) / 2
也就是说，原题目转化成了“找到nums当中有多少个子集A，满足sum(A) = (target + sum(nums)) / 2”
这子集问题其实也是一个背包问题，按照背包问题的套路，dp函数可以这样定义：
	dp[i][j] = x 表示当只在前i个数中选择时，若当前背包容量为j，那么可以有x种方法装满这个背包
定义好了dp数组后，就需要给出状态转移方程。在本题中选择较为简单，对num[i]只有选或不选两种。
若不选择num[i]，那么dp[i][j] = dp[i-1][j];
若选择num[i]，那么就需要看前i-1个物品有多少种方法能装满j - num[i-1]

*/

func subset(nums []int, sum int) int {
	//初始化一个dp table。根据dp数组的定义，最后返回的结果应该是dp[n][sum]
	n := len(nums)
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, sum+1))
	}
	//填充base case。dp[0][j] = 0,dp[i][0] = 1(什么都不装)
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j] //背包装不下，只能选择不装
			}
		}
	}

	return dp[n][sum]
}

func TargetSumWaysDp(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	//当sum的值小于target或者(target + sum)不能整除2时，无解
	if sum < target || (target+sum)%2 != 0 {
		return 0
	}
	return subset(nums, sum)
}
