package s771

import "testing"

func Test771(t *testing.T) {
	tt := []struct {
		jewels string
		stones string
		result int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, tc := range tt {
		r := numJewelsInStones(tc.jewels, tc.stones)
		if r != tc.result {
			t.Errorf("jewels = %v, stones = %v | should be %v got %v", tc.jewels, tc.stones, tc.result, r)
		}
	}
}
