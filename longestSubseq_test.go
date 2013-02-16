package algorithms

import (
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	testcases := []struct {
		seq    []int
		incSeq []int
	}{
		{[]int{5, 2, 8, 6, 3, 6, 9, 7}, []int{2, 3, 6, 7}},
		{[]int{9, 5, 2, 8, 6, 7}, []int{2, 6, 7}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1}},
		{[]int{2, 1, 2}, []int{1, 2}},
		// Van der Corput sequence
		{[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
			[]int{0, 2, 6, 9, 11, 15}},
	}

	for _, testcase := range testcases {
		actual := LongestIncreasingSubsequence(testcase.seq)
		if !intSliceEquals(actual, testcase.incSeq) {
			t.Errorf("Expected %v, got %v", testcase.incSeq, actual)
		}
	}
}

func intSliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
