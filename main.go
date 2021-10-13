package main

import (
	"Algo/dp"
	"fmt"
)

func testFib() {
	fmt.Println(dp.FibBase(6))
	fmt.Println(dp.FibMemo(6))
	fmt.Println(dp.FibDp(6))
	fmt.Println(dp.FibDpModify(6))
}

func testCoinChange() {
	fmt.Println(dp.CoinChangeRecursion(11,[]int{1,2,5}))
	fmt.Println(dp.CoinChangeIter(11,[]int{1,2,5}))
	fmt.Println(dp.CoinChangeRecModify(11,[]int{1,2,5}))
}

func main() {
	//testFib()
	testCoinChange()
}
