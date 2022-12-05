package excel_sheet_column_title

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Case 1",
			args{1},
			"A",
		},
		{
			"Case 2",
			args{28},
			"AB",
		},
		{
			"Case 3",
			args{701},
			"ZY",
		},
		{
			"Custom case 1",
			args{27},
			"AA",
		},
		{
			"Custom case 2",
			args{53},
			"BA",
		},
		{
			"Custom case 3",
			args{703},
			"AAA",
		},
		{
			"Custom case 4",
			args{1379},
			"BAA",
		},
		{
			"Custom case 5",
			args{18279},
			"AAAA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
