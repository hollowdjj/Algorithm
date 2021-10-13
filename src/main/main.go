package main

import (
	dp2 "Algo/src/dp"
	"fmt"
)

func testFib() {
	fmt.Println(dp2.FibBase(6))
	fmt.Println(dp2.FibMemo(6))
	fmt.Println(dp2.FibDp(6))
	fmt.Println(dp2.FibDpModify(6))
}

func testCoinChange() {
	fmt.Println(dp2.CoinChangeRecursion(11, []int{1, 2, 5}))
	fmt.Println(dp2.CoinChangeIter(11, []int{1, 2, 5}))
	fmt.Println(dp2.CoinChangeRecModify(11, []int{1, 2, 5}))
}

func main() {
	//testFib()
	testCoinChange()
}
