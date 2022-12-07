package valid_sudoku

// Constraints:
//   - board.length == 9
//   - board[i].length == 9
//   - board[i][j] is a digit 1-9 or '.'.

// N: height of the input board (fixed as 9)
// M: width of the input board (fixed as 9)
//
//   - Time Complexity: O(N*M)
//   - Space Complexity: O(N*M)
func isValidSudoku(board [][]byte) bool {
	cols, squares := [9][9]bool{}, [3][3][9]bool{}

	for r := range board {
		row := [9]bool{}

		for c := range board[r] {
			cell := board[r][c]
			if cell == '.' {
				continue
			}

			digitIndex := cell - '1' // '1' -> 0, '2' -> 1, ..., '9' -> 8

			if row[digitIndex] {
				return false
			}
			row[digitIndex] = true

			if cols[c][digitIndex] {
				return false
			}
			cols[c][digitIndex] = true

			if squares[r/3][c/3][digitIndex] {
				return false
			}
			squares[r/3][c/3][digitIndex] = true
		}
	}

	return true
}
