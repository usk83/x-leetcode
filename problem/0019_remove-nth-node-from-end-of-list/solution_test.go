package remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"

	"github.com/usk83/x-leetcode/internal/ds/list"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *list.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *list.ListNode
	}{
		{
			"Case 1",
			args{
				list.MustNewFromString("[1,2,3,4,5]"),
				2,
			},
			list.MustNewFromString("[1,2,3,5]"),
		},
		{
			"Case 2",
			args{
				list.MustNewFromString("[1]"),
				1,
			},
			list.MustNewFromString("[]"),
		},
		{
			"Case 3",
			args{
				list.MustNewFromString("[1,2]"),
				1,
			},
			list.MustNewFromString("[1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
