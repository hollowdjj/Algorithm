package dp

/*
最长公共子序列问题：给定两个字符串s1,s2，找出他们两的最长公共子序列，返回这个子序列的长度，例如：
输入：s1="zabcde" s2="acez"
输出：3
解释：最长公共子序列为“ace”
*/

/*
Lcs
dp数组：dp[i][j]表示字符串s1[0...i]与s2[0...j]的最长公共子序列。
状态转移方程：
                                  0                  i = 0或j = 0
      dp[i][j]=	 dp[i-1][j-1] + 1		             s1[i] == s2[j]
                 max{dp[i-1][j],dp[i][j-1]}          s1[i] != s2[j]
*/
func Lcs(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	//初始化一个dp table
	var dp [][]int
	for i := 0; i <= n1; i++ {
		dp = append(dp, make([]int, n2+1))
	}
	//选择
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] != s2[j-1] {
				dp[i][j] = maxOfTwo(dp[i-1][j], dp[i][j-1])
			} else if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return dp[n1][n2]
}

//LcsCompressed dp[i][j]的值只和dp[i-1][j-1] dp[i-1][j] dp[i][j-1]有关，因此，可以进一步优化
func LcsCompressed(s1, s2 string) int {
	//压缩成一个一维dp table 空间复杂度为O(min{len(s1),len(s2)})
	n1, n2 := len(s1), len(s2)
	var dp []int
	if n1 <= n2 {
		dp = make([]int, n1+1)
		s1, s2 = s2, s1
		n1, n2 = n2, n1
	} else {
		dp = make([]int, n2+1)
	}
	//选择
	upLeft := 0
	for i := 1; i <= n1; i++ {
		upLeft = dp[0]
		for j := 1; j <= n2; j++ {
			temp := dp[j]
			if s1[i-1] != s2[j-1] {
				dp[j] = maxOfTwo(dp[j-1], dp[j])
			} else if s1[i-1] == s2[j-1] {
				dp[j] = upLeft + 1
			}
			upLeft = temp
		}
	}

	return dp[n2]
}

func maxOfTwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minOfTwo(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
