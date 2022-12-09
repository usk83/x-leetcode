package insert_interval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Case 1",
			args{
				[][]int{
					{1, 3},
					{6, 9},
				},
				[]int{2, 5},
			},
			[][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			"Case 2",
			args{
				[][]int{
					{1, 2},
					{3, 5},
					{6, 7},
					{8, 10},
					{12, 16},
				},
				[]int{4, 8},
			},
			[][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
		{
			"Custom case 1",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{0, 5},
			},
			[][]int{
				{0, 5},
				{10, 15},
				{20, 25},
				{30, 35},
				{40, 45},
			},
		},
		{
			"Custom case 2",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{5, 10},
			},
			[][]int{
				{5, 15},
				{20, 25},
				{30, 35},
				{40, 45},
			},
		},
		{
			"Custom case 3",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{5, 30},
			},
			[][]int{
				{5, 35},
				{40, 45},
			},
		},
		{
			"Custom case 4",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{25, 27},
			},
			[][]int{
				{10, 15},
				{20, 27},
				{30, 35},
				{40, 45},
			},
		},
		{
			"Custom case 5",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{28, 30},
			},
			[][]int{
				{10, 15},
				{20, 25},
				{28, 35},
				{40, 45},
			},
		},
		{
			"Custom case 6",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{26, 29},
			},
			[][]int{
				{10, 15},
				{20, 25},
				{26, 29},
				{30, 35},
				{40, 45},
			},
		},
		{
			"Custom case 7",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{32, 44},
			},
			[][]int{
				{10, 15},
				{20, 25},
				{30, 45},
			},
		},
		{
			"Custom case 8",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{45, 50},
			},
			[][]int{
				{10, 15},
				{20, 25},
				{30, 35},
				{40, 50},
			},
		},
		{
			"Custom case 9",
			args{
				[][]int{
					{10, 15},
					{20, 25},
					{30, 35},
					{40, 45},
				},
				[]int{50, 55},
			},
			[][]int{
				{10, 15},
				{20, 25},
				{30, 35},
				{40, 45},
				{50, 55},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
