package dp

/*
编辑距离：给定两个字符串s1和s2，计算出将s1转换成s2所使用的最少操作数。你可以对一个字符串进行如下三种操作：
1. 插入一个字符 2. 删除一个字符 3. 替换一个字符。例如：
输入：s1="horse", s2="ros"
输出: 3
解释：
horse -> rorse(将'h'替换为'r')
rorse -> rose(删除'r')
rose  -> ros(删除's')

分析：
需要记住的是，解决两个字符串的动态规划问题，一般都是用两个指针i,j分别指向两个字符串的最后，然后一步步往前
走，缩小问题的规模。从而算法的base case有两个：
1. 当s1走完了，s2还没有走完，即i==-1且j>=0时，只需要把s2中剩下的字符插入到s1前面即可
2. 当s2走完了，s1还没有走完，即i>=0且j==-1时，只能用删除操作把s1剩下的字符全部删除
*/

/*
MinDistanceRecursion
递归方式解决。dp函数定义为：dp(i,j)返回s1[0,1...i]转换到s2[0,1...j]的最小步骤数
*/
func MinDistanceRecursion(s1, s2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		//base case
		if i == -1 {
			return j + 1 //此时只需要把s2[0,..j]插入到s1前面就行，故需要j+1步
		}
		if j == -1 {
			return i + 1 //此时只需要把s1[0,..i]这些字符全部删除即可，故需要i+1步
		}
		//两字符相等，不需要进行任何操作
		if s1[i] == s2[j] {
			return dp(i-1, j-1)
		} else { //两字符不相等，那么就需要比较插入、删除或替换这三种操作，哪一个步骤最少
			return minOfThree(dp(i, j-1)+1, //插入。s[j]已经匹配了，那么j往前移一格
				dp(i-1, j)+1,   //删除。删除s[i]，那么i往前移一格，j不变
				dp(i-1, j-1)+1) //替换。意味着s1[i] == s2[j]，故i,j均往前移动一格
		}
	}

	return dp(len(s1)-1, len(s2)-1)
}

/*
MinDistanceDp
使用动态规划优化上述递归过程。动态规划中的优化无法就是消除重叠子问题，然后进行状态压缩。分析三个选择可知，
我们在计算dp(i-1,j)时，又会去计算dp(i-1,j-1)，计算dp(i,j-1)时，也会去计算dp(i-1,j-1)。因此，存在大量
重叠子问题。定义dp数组dp[i][j]表示s1[0...i]转换成s2[0..j]的最小步骤数。从而可以有状态转移方程：
			dp[i][0] == i; dp[0][j] == j
dp[i][j] =  min{dp[i-1][j],dp[i][j-1],dp[i-1][j-1]} + 1   s1[i-1] != s2[j-1]
            dp[i-1][j-1]                                  s1[i-1] == s2[j-1]
此题很明显的体现出了，递归程序的代码可阅读性要好一些，这里dp table的自底向上不太好理解。
*/
func MinDistanceDp(s1, s2 string) int {
	/*
		创建并初始化一个dp table，最后返回的结果为dp[n1][n2]。这样设计的原因是因为，base case为i和j分别为
		-1的情况，但数组的索引最少只能到0，因此需要左移一格。
	*/
	n1, n2 := len(s1), len(s2)
	var dp [][]int
	for i := 0; i <= n1; i++ {
		dp = append(dp, make([]int, n2+1))
	}
	//base case dp[0][j] == j; dp[i][0] == i(由于左移了一格，0代表-1，j就表示原来的j+1)
	for i := 1; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n2; j++ {
		dp[0][j] = j
	}
	//选择
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] { //当s1[i-1] == s2[j-1]时，不需要做任何操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minOfThree(dp[i-1][j]+1, //删除。删除s[i]，那么i往前移一格，j不变
					dp[i-1][j-1]+1, //替换。意味着s1[i] == s2[j]，故i,j均往前移动一格
					dp[i][j-1]+1)   //插入。s[j]已经匹配了，那么j往前移一格
			}
		}
	}

	return dp[n1][n2]
}

func MinDistanceDpCompressed(s1, s2 string) int {
	//二维dp table状态压缩后的一个一维dp table。空间复杂度为O(min{len(s1),len(s2)})
	n1, n2 := len(s1), len(s2)
	var dp []int
	if n1 < n2 {
		dp = make([]int, n1+1)
		n1, n2 = n2, n1 //交换一下方便操作
		s1, s2 = s2, s1
	} else {
		dp = make([]int, n2+1)
	}
	//base case
	for i := 1; i < len(dp); i++ {
		dp[i] = i
	}
	//选择
	for i := 1; i <= n1; i++ {
		dp[0] = i
		left, topLeft, top := i, dp[0], dp[1]
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[j] = topLeft
			} else {
				dp[j] = minOfThree(topLeft+1, top+1, left+1)
			}

			if j == n2 {
				break
			}
			topLeft = top
			top = dp[j+1]
			left = dp[j]

		}
	}

	return dp[len(dp)-1]
}
