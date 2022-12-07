package letter_combinations_of_a_phone_number

// Constraints:
//   - 0 <= digits.length <= 4
//   - digits[i] is a digit in the range ['2', '9'].

var (
	letterCombinations = letterCombinationsSimple
	_                  = letterCombinationsSpacePreallocated
)

var digitsLettersMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// N: length of the input string digits
//
//   - Time Complexity: O(4^N)?
//   - Space Complexity: O(4^N)?
func letterCombinationsSimple(digits string) []string {
	combs := []string{}

	if len(digits) == 0 {
		return combs
	}

	var dfs func([]byte, int)
	dfs = func(bs []byte, i int) {
		if i == len(digits) {
			combs = append(combs, string(bs))
			return
		}

		for _, b := range digitsLettersMap[digits[i]] {
			dfs(append(bs, b), i+1)
		}
	}
	dfs(make([]byte, 0, len(digits)), 0)

	return combs
}

// N: length of the input string digits
//
//   - Time Complexity: O(4^N)?
//   - Space Complexity: O(4^N)?
func letterCombinationsSpacePreallocated(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combsSize := 1
	for i := range digits {
		combsSize *= len(digitsLettersMap[digits[i]])
	}

	combs := make([]string, 0, combsSize)
	bytesHolder := make([]byte, len(digits))

	var dfs func(int)
	dfs = func(i int) {
		if i == len(digits) {
			combs = append(combs, string(bytesHolder))
			return
		}

		for _, b := range digitsLettersMap[digits[i]] {
			bytesHolder[i] = b
			dfs(i + 1)
		}
	}
	dfs(0)

	return combs
}
