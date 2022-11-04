package sqrtx

import (
	"math"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// default test cases
		{"Case 1", args{4}, 2},
		{"Case 2", args{8}, 2},

		// custom test cases
		{"Case 3: Smallest number", args{0}, 0},
		{"Case 4: Largest number", args{int(math.Exp2(31)) - 1}, 46340},
		{"Case 5: Square number", args{121}, 11},
		{"Case 6: Small number", args{1}, 1},
		{"Case 7: Small number", args{2}, 1},
		{"Case 8: Small number", args{3}, 1},
		{"Case 9: Random number", args{12471412}, 3531},
		{"Case 10: Random number", args{876976134}, 29613},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
