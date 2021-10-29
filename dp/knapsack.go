package dp

/*
0-1背包问题：有一个可以装载重量为W的背包和N个物体。每个物体有重量和价值两个属性，其中第i个物体的重量为wt[i]
价值为val[i]。现用这个背包装物体，最多能装的价值是多少。
*/

/*
Knapsack
dp数组的定义为：dp[i][j]表示，若只在前i个物体中选择，且当前背包容量为j且时，背包中能装的最大价值。返回dp[N][W]
显然，对第i个物体(i从1开始)，它的选择只有装入背包和不装入背包两种：
1. 若第i个物体不装入背包，那么dp[i][j] = dp[i-1][j]。
2. 若第i个物体装入背包，那么dp[i][j] = dp[i-1][j-wt[i-1]] +val[i-1] (wt[i-1]和val[i-1]分别表示第i个物体的重量和质量)
这里解释一下为什么是dp[i-1][j-wt[i-1]]。因为，若选择将第i个物体装入背包，而当前背包的容量又是w，那么背包剩下的
容量就是j-wt[i-1]，显然我们需要寻求在只选择前i-1个物体且剩余容量为j-wt[i-1]的最大价值
*/
func Knapsack(n, w int, wt []int, val []int) int {
	//初始化dp table
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, w+1))
	}
	//选择
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] >= 0 {
				dp[i][j] = maxOfTwo(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][w]
}
