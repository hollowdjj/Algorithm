package dp

/*
动态规划问题求解的关键在于列出状态转移方程。要想正确列出状态状态方程需要明确以下4点内容：
1. 明确base case  			 即最基础的情况
2. 明确[状态]      			 所谓状态就是原问题和子问题中会不断变化的量
3. 明确[选择]                 即导致状态变化的行为
4. 定义dp数组或函数的意义      不同的定义方法，会导致递归dp或循环dp
一旦理清了以上四点后，就可以套下面这个框架进行求解：
dp[0][0][...] = base
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 求最值(选择1，选择2...)
*/

/*!
斐波拉契数列f(n)的计算方式为：f(0) = 0  f(1) = 1; 当n大于等于2时，f(n) = f(n-1) + f(n-2)
因此，很显然可以使用递归的方式进行求解。然而，这种递归方式存在大量的重复计算。例如，当我们要计
算f(20)时，我们首先需要计算f(18)和f(19)，而计算f(19)又需要计算f(18)和f(17)，也就是说同一个
值f(n)会被多次计算，这无疑会显著增加算法的时间复杂度。递归算法的时间复杂度为“子问题个数乘以计
算一个子问题所需要的时间”。在FibBase函数中，子问题的个数即为递归树的节点个数，而该递归树为一
颗二叉树，其节点个数是指数级别的，即O(2^n)。每一个子问题都是加法运算，故时间复杂度为O(1)，因
此总的时间复杂度为O(2^n)
*/

func FibBase(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return FibBase(n-1) + FibBase(n-2)
}

/*!
在FibBase中存在子问题被多次重复计算的问题。因此，我们很容易想到使用一个“备忘录”来记录每一个子问题
的计算结果，从而每个子问题只会计算一次，时间复杂度由O(2^n)降为O(n)。这种递归计算方法被称为“自顶向下”，即
将一个较大规模的问题，例如f(20)向下逐渐分解，直到base case，与动态规划的“自底向上”刚好相反。
*/

func helper(n int, m map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}
	//如果子问题已经计算过了，那么直接返回
	if val, ok := m[n]; ok {
		return val
	}
	//计算子问题
	m[n] = helper(n-1, m) + helper(n-2, m)
	return m[n]
}

func FibMemo(n int) int {
	memo := make(map[int]int, n+1) //用哈希表实现一个备忘录。key为n,value为f(n)
	return helper(n, memo)
}

/*!
动态规划是一种“自底向上”的方法。也就是说，在计算f(20)时，会首先从f(0),f(1)....计算起，一直计算到f(20)。因此，
动态规划一般使用的是循环而非递归。在动态规划中，最重要的是找到“状态转移方程”。在这里，f(n)可以视作一个状态n，而
f(n) = f(n-1) + f(n-2)，即状态n等于上一个状态和上上个状态之和。FibDp的时间复杂度和空间复杂度均为O(n)。

简单分析一下，不难发现，在斐波拉契数列问题中：
base case为：f(n) == n (n == 0 || n == 1)
状态为：     函数值即f(n)
选择：       即函数f(n) = f(n-1) + f(n-2)
dp函数定义：  dp(n) 输入数字n，返回斐波拉契数列f(n)的值
*/

func FibDp(n int) int {
	table := make([]int, n+1) //dp table
	table[0], table[1] = 0, 1 //base case

	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2] //选择
	}

	return table[n]
}

/*!
在FibDp中，不难发现，计算当前状态的值，只需要前两个状态。因此，可以进一步将空间复杂度优化至O(1)
*/

func FibDpModify(n int) int {
	prevOfPrev, prev := 0, 1
	for i := 2; i <= n; i++ {
		temp := prev
		prev = prev + prevOfPrev
		prevOfPrev = temp
	}

	return prev
}
