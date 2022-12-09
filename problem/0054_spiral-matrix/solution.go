package spiral_matrix

// Constraints:
//   - m == matrix.length
//   - n == matrix[i].length
//   - 1 <= m, n <= 10
//   - -100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) []int {
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	m, n := len(matrix), len(matrix[0])
	size := [2]int{n, m}
	orderedElems := make([]int, 0, m*n)

	p := [2]int{0, -1}
	for i := 0; size[i&1] > 0; i++ {
		for j := 0; j < size[i&1]; j++ {
			p[0] += dir[i%4][0]
			p[1] += dir[i%4][1]
			orderedElems = append(orderedElems, matrix[p[0]][p[1]])
		}
		size[i&1^1]--
	}

	return orderedElems
}
