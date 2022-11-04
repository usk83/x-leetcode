package search_a_2d_matrix

import "sort"

// Constraints:
//   - m == matrix.length
//   - n == matrix[i].length
//   - 1 <= m, n <= 100
//   - -10^4 <= matrix[i][j], target <= 10^4

// Do binary search twice
//   - Time Complexity: O(logN+logM)
//   - Space Complexity: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] {
		return false
	}

	mi := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1

	ni := sort.Search(len(matrix[mi]), func(i int) bool {
		return matrix[mi][i] > target
	}) - 1

	return matrix[mi][ni] == target
}
