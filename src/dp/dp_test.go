package dp

import "testing"

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
