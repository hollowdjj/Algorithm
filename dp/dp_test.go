package dp

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {

	fmt.Println(CoinChangeRecursion(11, []int{1, 2, 5}))
	fmt.Println(CoinChangeIter(11, []int{1, 2, 5}))
	fmt.Println(CoinChangeRecModify(11, []int{1, 2, 5}))
}

func TestMaxEnvelopes(t *testing.T) {
	testData := []Envelope{{Width: 5, Height: 4},
		{Width: 6, Height: 4},
		{Width: 6, Height: 7},
		{Width: 2, Height: 3}}
	//save := testData
	//resData := []Envelope{{2,3},{5,4},{6,7},{6,4}}
	//MaxEnvelopes(testData)
	////if envelopeSliceEqual(testData,resData) == false {
	////	t.Errorf("testData:\n%v\nwanting:\n%v\nbut get:\n%v\n",save,resData,testData)
	////}
	res := MaxEnvelopes(testData)
	if res != 3 {
		t.Errorf("The result for:\n%v\nshould be %d instead of %d", testData, 3, res)
	}
}

func TestMaxSubArraySum(t *testing.T) {
	testData := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := MaxSubArraySum(testData)
	res1 := MaxSubArraySumCompressed(testData)
	if res != 6 {
		t.Errorf("The result for %v should be %d, but got %d", testData, 6, res)
	}
	if res1 != 6 {
		t.Errorf("The result for %v should be %d, but got %d", testData, 6, res1)
	}
}

func TestLcs(t *testing.T) {
	s1, s2 := "abcde", "ace"
	res := Lcs(s1, s2)
	res1 := LcsCompressed(s1, s2)
	if res != 3 && res1 != 3 {
		t.Errorf("Something is wrong\n")
	}
}
