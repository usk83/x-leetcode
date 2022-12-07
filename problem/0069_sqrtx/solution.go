package sqrtx

import (
	"math/bits"
	"sort"
)

var (
	mySqrt = manualBinarySearch
	_      = stdlibSortSearch
	_      = newtonsMethod
	_      = bitwiseOperations
)

// N: input number x
//
//   - Time Complexity: O(logN)
//   - Space Complexity: O(1)
func manualBinarySearch(x int) int {
	// Find the smallest positive integer(let's say "the target number") that exceeds `x` when squared.

	l, r := 1, x+1 // `r` should be `x+1` since the target number does not exist in `[1, x]` when `x <= 1`

	for l < r {
		if c := (l + r) >> 1; c*c <= x { // no need to consider overflow on 64-bit system (since `0 <= x <= 2^31 - 1`)
			l = c + 1
		} else {
			r = c
		}
	}

	// Then return `the target number - 1`.
	return l - 1
}

// N: input number x
//
//   - Time Complexity: O(logN)
//   - Space Complexity: O(1)
func stdlibSortSearch(x int) int {
	// Find the smallest positive integer that exceeds `x` when squared.
	// Then return that number - 1.
	return sort.Search(x+1, func(i int) bool {
		return i*i > x
	}) - 1
}

// N: input number x
//
//   - Time Complexity: O(logN)?
//   - Space Complexity: O(1)
//
// This code was taken from [Golang, Integer Newton]
//
// [Golang, Integer Newton]: https://leetcode.com/problems/sqrtx/solutions/393944/golang-integer-newton/
func newtonsMethod(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// L: minimum number of bits required to represent x
//
//   - Time Complexity: O(L)
//   - Space Complexity: O(1)
//
// This code was written with reference to [Golang fast and simple solution, binary-search, bit manipulation]
//
// [Golang fast and simple solution, binary-search, bit manipulation]: https://leetcode.com/problems/sqrtx/solutions/2718066/golang-fast-and-simple-solution-binary-search-bit-manipulation/
func bitwiseOperations(x int) int {
	xLen := bits.Len(uint(x))
	maxLen := xLen>>1 + xLen&1 // halve and round up

	var num int
	for pos := maxLen - 1; pos >= 0; pos-- {
		// set a bit at the current position and see if the temporary number is not too large
		if tmp := num | 1<<pos; tmp*tmp <= x {
			num = tmp
		}
	}

	return num
}
