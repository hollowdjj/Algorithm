package dp

import "math"

/*
凑零钱问题：给定k种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，
问最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1。
分析：
base case: amount为0时，需要0枚硬币
状态：      是指原问题和子问题中会变化的量。在这里即为目标金额
选择：      导致状态变化的行为。只有选择硬币才会让目标金额产生变化，故所有硬币的面值就是选择
dp函数定义： dp(n)，输入一个目标金额n，返回最少需要m枚硬币凑出这个金额(递归)
状态转移方程为：
         -1  												(n < 0)
dp(n) =   0   												(n = 0)
         min(dp(n - coin) + 1,current_result | coin∈coins) (n > 0)
递归树的节点数量为O(n^k)，子问题的时间复杂度为O(k)，故总的时间复杂度为O(k * n^k)
*/

func CoinChangeRecursion(n int,coins []int) int {
	//base case
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}

	res := math.MaxInt
	for _,val := range coins {
		newRes := CoinChangeRecursion(n-val,coins)     //子问题
		if newRes == -1 {
			continue
		}
		if newRes + 1 < res {               		   //选择
			res = newRes + 1
		}
	}

	return res
}

/*
在函数CoinChangeRecursion中存在重复子问题，可采取“备忘录”进一步优化。修改后，子问题个数为O(n)，算法时间
复杂度降低为O(kn)，空间复杂度增加为O(n)
*/
var memo = make(map[int]int)    		          //key为金额，value为凑出这个金额至少需要的硬币数量
func CoinChangeRecModify(n int,coins []int) int {
	//base case
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	//递归
	res := math.MaxInt
	for _,val := range coins {
		if n - val < 0 {
			continue
		}
		//如果子问题已经计算过了，那么就跳过
		newRes := 0
		if num,ok := memo[n-val];ok {
			newRes = num
		} else {
			newRes = CoinChangeRecModify(n-val,coins)
			memo[n-val] = newRes
		}
		//执行选择以更新res
		if newRes == -1 {
			continue
		}
		if newRes + 1 < res {
			res = newRes + 1
		}
	}

	return res
}

/*
dp数组定义：dp[i]，当目标金额为i时，至少需要dp[i]枚硬币
状态转移方程：
         -1  					  (i < 0)
dp[i] =   0   					  (i = 0)
          min(dp[i],dp[i-ck] + 1) (i > 0)
时间复杂度为O(kn)，空间复杂度为O(n)
*/

func CoinChangeIter(n int,coins []int) int {
	//dp table
	dp := make([]int,n+1)
	for i:=0;i<len(dp);i++ {
		dp[i] = n + 1
	}
	//base case
	dp[0] = 0
	//迭代。对dp[i]需要尝试选择所有面值的硬币
	for i:=0;i < len(dp);i++ {
		for _,val := range coins {
			if i - val < 0 {
				continue
			}
			//选择
			if dp[i] > dp[i - val] + 1 {
				dp[i] = dp[i - val] + 1
			}
		}
	}

	if dp[n] == n + 1 {
		return -1
	}

	return dp[n]
}


