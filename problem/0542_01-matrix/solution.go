package _01_matrix

import "math"

var (
	_            = updateMatrixDFS
	updateMatrix = updateMatrixBFS
)

// Constraints:
//   - m == mat.length
//   - n == mat[i].length
//   - 1 <= m, n <= 10^4
//   - 1 <= m * n <= 10^4
//   - mat[i][j] is either 0 or 1.
//   - There is at least one 0 in mat.

func updateMatrixDFS(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dir := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	distMap := make([][]int, m)
	for y := range distMap {
		distMap[y] = make([]int, n)
		for x := range distMap[y] {
			distMap[y][x] = math.MaxInt
		}
	}

	var dfs func(int, int, int)
	dfs = func(y, x, count int) {
		if y < 0 || y >= m || x < 0 || x >= n ||
			mat[y][x] == 0 ||
			distMap[y][x] <= count {
			return
		}

		distMap[y][x] = count

		for _, d := range dir {
			dfs(y+d[0], x+d[1], count+1)
		}
	}

	for y := range mat {
		for x := range mat[y] {
			if mat[y][x] == 1 {
				continue
			}

			distMap[y][x] = 0

			for _, d := range dir {
				dfs(y+d[0], x+d[1], 1)
			}
		}
	}

	return distMap
}

func updateMatrixBFS(mat [][]int) [][]int {
	dir := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n := len(mat), len(mat[0])

	distMap := make([][]int, m)
	for y := range distMap {
		distMap[y] = make([]int, n)
		for x := range distMap[y] {
			distMap[y][x] = math.MaxInt
		}
	}

	var queue [][2]int
	for y := range mat {
		for x := range mat[y] {
			if mat[y][x] == 0 {
				distMap[y][x] = 0
				queue = append(queue, [2]int{y, x})
			}
		}
	}

	var count int
	for len(queue) != 0 {
		count++

		var nextQueue [][2]int
		for _, coord := range queue {
			for _, d := range dir {
				y, x := coord[0]+d[0], coord[1]+d[1]

				if y >= 0 && y < m && x >= 0 && x < n && distMap[y][x] > count {
					distMap[y][x] = count
					nextQueue = append(nextQueue, [2]int{y, x})
				}
			}
		}
		queue = nextQueue
	}

	return distMap
}
