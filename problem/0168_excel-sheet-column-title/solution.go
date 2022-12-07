package excel_sheet_column_title

// Constraints:
//   - 1 <= columnNumber <= 2^31 - 1

// N: input number columnNumber
//
//   - Time Complexity: O(log_26 N)
//   - Space Complexity: O(log_26 N)
func convertToTitle(columnNumber int) string {
	const base = 26

	var titleChars []byte

	for columnNumber > 0 {
		columnNumber--
		titleChars = append(titleChars, 'A'+byte(columnNumber%base))
		columnNumber /= base
	}

	for i := 0; i < len(titleChars)>>1; i++ {
		titleChars[i], titleChars[len(titleChars)-1-i] = titleChars[len(titleChars)-1-i], titleChars[i]
	}

	return string(titleChars)
}
