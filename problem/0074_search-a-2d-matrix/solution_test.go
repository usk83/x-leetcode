package search_a_2d_matrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// default test cases
		{"Case 1", args{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
		}, true},
		{"Case 2", args{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			13,
		}, false},

		// custom test cases
		{"Case 3: Smallest matrix", args{
			[][]int{
				{10},
			},
			12,
		}, false},
		{"Case 4: Target is smaller than any numbers in the matrix", args{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			0,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
