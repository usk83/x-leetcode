package insert_interval

import (
	"sort"
)

// Constraints:
//   - 0 <= intervals.length <= 10^4
//   - intervals[i].length == 2
//   - 0 <= start_i <= end_i <= 10^5
//   - intervals is sorted by start_i in ascending order.
//   - newInterval.length == 2
//   - 0 <= start <= end <= 10^5

func insert(intervals [][]int, newInterval []int) [][]int {
	left := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][1] >= newInterval[0]
	})

	right := sort.Search(len(intervals)-left, func(i int) bool {
		return intervals[i+left][0] > newInterval[1]
	}) + left

	var newIntervals [][]int

	newIntervals = append(newIntervals, intervals[:left]...)

	if left == right {
		newIntervals = append(newIntervals, newInterval)
	} else {
		newIntervals = append(newIntervals, []int{
			min(intervals[left][0], newInterval[0]),
			max(intervals[right-1][1], newInterval[1]),
		})
	}

	newIntervals = append(newIntervals, intervals[right:]...)

	return newIntervals
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
