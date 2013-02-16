// Package algorithms is my personal playground for learning new
// algorithms by implementing them. Since the purpose is educational,
// the implementations will generally not be optimized for
// performance.
//
// Thomas Kappler <tkappler@gmail.com>
package algorithms

import (
	//"fmt"
	"sort"
)

// Given an integer sequence seq, unsorted, find the longest
// subsequence, not necessarily consecutive, whose elements are sorted
// increasingly. There may be more than one, in this case this
// implementation chooses arbitrarily.
//
// The algorithm is in O(n log n) and is from
// http://en.wikipedia.org/wiki/Longest_increasing_subsequence,
// 2013-02-16.
func LongestIncreasingSubsequence(seq []int) (incSeq []int) {
	lseq := len(seq)

	// A sequence [] or [n] is its own longest subsequence.
	if lseq < 2 {
		return seq
	}

	// M[j] stores the index i in seq so that there is a subsequence
	// of length j ending at seq[i].
	M := make([]int, 1, lseq)
	M[0] = 0
	// The length of the longest subsequence found so far.
	L := 1
	// The sequence predecessors for each element referenced by M.
	P := make([]int, lseq)

	// Print the longest subsequence from the book-keeping variables
	// above. M[len(M)-1] is the final element of the sequence, since
	// a sequence ending at M[j] is of length[j]. Then we fill in the
	// rest of the sequence backwards, looking up each element in P.
	longestSubsequence := func() []int {
		l := len(M)
		lseq := make([]int, l)
		idx := M[l-1]
		for i := l - 1; i >= 0; i-- {
			lseq[i] = seq[idx]
			idx = P[idx]
		}
		return lseq
	}

	// The core of the algorithm. Look at each element in seq and
	// determine how it can extend a subsequence.
	for i := 1; i < lseq; i++ {
		lenSubseq := len(M)

		cur := seq[i]
		// Try to find a sequence predecessor of cur.
		j := sort.Search(lenSubseq, func(k int) bool {
			return seq[M[k]] >= cur
		})

		if j > 0 && seq[M[j-1]] < cur {
			// We have one, record it.
			P[i] = M[j-1]
		}

		if j == L || cur < seq[M[j]] {
			// We extended the current longest subsequence.
			if j >= lenSubseq {
				M = append(M, i)
			} else {
				M[j] = i
			}

			if j+1 > L {
				L++
			}
		}
	}

	return longestSubsequence()
}
