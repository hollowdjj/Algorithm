package dp

import "math"

/*
下降路径最小和问题。即初始时，可以站在矩阵matrix第一行中的任意一个元素，需要下降到最后一行。每次下降
可以向正下方、左下方和右下方移动一格，也就是说，matrix[i,j]可以下降到matrix[i,j+1],matrix[i-1,j+1]
或matrix[i+1,j+1]。相反，能够下降到matrix[i,j]的只有matrix[i,j-1]，matrix[i+1,j-1]或matrix[i-1,j-1]
           2   1   3
           6   5   4
           7   8   9      下降路径最小和为1 + 5 + 7
*/

/*
dp函数定义：	 dp[i,j]表示下降到matrix[i,j]的最小路径和
             dp[i,j] = i 									  				 i = 0
状态转移方程: dp[i,j] = min{dp[i,j-1],dp[i+1,j-1],dp[i-1,j-1]} + matrix[i,j]  i,j>=1
空间复杂度O(n^2) 时间复杂度O(n^2)
*/

func MinFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	//创建一个二维dp table，长度为dp[n][n]
	var dp [][]int
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
		dp[0][i] = matrix[0][i]
	}
	//选择
	for i := 1; i < n; i++ {
		//处理j=0的情况
		min := minOfThree(dp[i-1][0], math.MaxInt, dp[i-1][1])
		dp[i][0] = min + matrix[i][0]
		for j := 1; j < n-1; j++ {
			min = minOfThree(dp[i-1][j], dp[i-1][j-1], dp[i-1][j+1])
			dp[i][j] = min + matrix[i][j]
		}
		//处理j=n-1的情况
		min = minOfThree(dp[i-1][n-1], dp[i-1][n-2], math.MaxInt)
		dp[i][n-1] = min + matrix[i][n-1]
	}
	//遍历dp table的最后一行，找到最小和
	res := dp[n-1][0]
	for j := 1; j < n; j++ {
		if res > dp[n-1][j] {
			res = dp[n-1][j]
		}
	}

	return res
}

/*
从状态转移方程: dp[i,j] = min{dp[i,j-1],dp[i+1,j-1],dp[i-1,j-1]} + matrix[i,j]  i,j>=1
可以看出，求dp[i,j]只需要前面三个状态，那么我们真的需要一个n*n的二维数组dp吗？答案是否定的，可
以通过状态压缩，将这个二维数组映射到一个一维数组，从而空间复杂度由O(n^2)降到了O(n)
 2  1  3
 0  0  0
 0  0  0
    ↓
 2  1  3
*/

func MinFallingPathSumCompressed(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		/*
			当处理到matrix的第i行时，dp数组表示的是matrix第i-1行中各个元素的下降路径最小和
			同时，在处理第i行时，需要更新dp数组，故需要使用临时变量保存那些会被更新的值
		*/

		//处理j=0的情况
		temp := dp[0]
		min := minOfThree(dp[0], dp[1], math.MaxInt)
		dp[0] = min + matrix[i][0]
		for j := 1; j < n-1; j++ {
			min = minOfThree(temp, dp[j], dp[j+1])
			temp = dp[j]
			dp[j] = min + matrix[i][j]
		}
		//处理j=n-1的情况
		min = minOfThree(temp, dp[n-1], math.MaxInt)
		dp[n-1] = min + matrix[i][n-1]
	}

	res := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] < res {
			res = dp[i]
		}
	}

	return res
}

func minOfThree(a, b, c int) int {
	min := a
	if b <= min {
		min = b
	}
	if c <= min {
		min = c
	}

	return min
}
