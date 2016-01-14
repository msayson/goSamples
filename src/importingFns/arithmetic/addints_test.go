package arithmetic

import "testing"

func TestAddInts(t *testing.T) {
	testCases := []struct {
		inA, inB, expect int
	}{
		{-5, 4, -1},
		{15, 4, 19},
		{0, 4, 4},
	}
	for _, test := range testCases {
		sum := AddInts(test.inA, test.inB)
		if sum != test.expect {
			t.Errorf("AddInts(%d, %d) == %d, expected %d", test.inA, test.inB, sum, test.expect)
		}
	}
}
