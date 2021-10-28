package main

import (
	dp "Algo/dp"
	"fmt"
)

func testFib() {
	fmt.Println(dp.FibBase(6))
	fmt.Println(dp.FibMemo(6))
	fmt.Println(dp.FibDp(6))
	fmt.Println(dp.FibDpModify(6))
}

func testCoinChange() {
	fmt.Println(dp.CoinChangeRecursion(11, []int{1, 2, 5}))
	fmt.Println(dp.CoinChangeIter(11, []int{1, 2, 5}))
	fmt.Println(dp.CoinChangeRecModify(11, []int{1, 2, 5}))
}

func testLis() {
	fmt.Println(dp.Lis([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func testMinFallingPath() {
	fmt.Println(dp.MinFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
	fmt.Println(dp.MinFallingPathSumCompressed([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}

func testTargetSumWays() {
	fmt.Println(dp.TargetSumWaysBackTrack([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(dp.TargetSumWaysDp([]int{1, 1, 1, 1, 1}, 3))
}

func testMinDistance() {
	fmt.Println(dp.MinDistanceRecursion("rad", "apple"))
	fmt.Println(dp.MinDistanceDp("rad", "apple"))
	fmt.Println(dp.MinDistanceDpCompressed("rad", "apple"))
}

func testMaxEnvelopes() {
	data := []dp.Envelope{{Width: 5, Height: 4},
		{Width: 6, Height: 4},
		{Width: 6, Height: 7},
		{Width: 2, Height: 3}}
	fmt.Println(dp.MaxEnvelopes(data))
}

func testMaxSubArraySum() {
	testData := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(dp.MaxSubArraySum(testData))
}

func testLcs() {
	s1, s2 := "bl", "yby"
	fmt.Println(dp.Lcs(s1, s2))
	fmt.Println(dp.LcsCompressed(s1, s2))
}

func main() {
	//testFib()
	//testCoinChange()
	//testLis()
	//testMinFallingPath()
	//testTargetSumWays()
	//testMinDistance()
	//testMaxEnvelopes()
	//testMaxSubArraySum()
	testLcs()
}
