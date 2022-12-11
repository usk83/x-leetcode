package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"

	"github.com/usk83/x-leetcode/internal/ds/tree"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Case 1",
			args{tree.MustNewFromString("[3,9,20,null,null,15,7]")},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			"Case 2",
			args{tree.MustNewFromString("[1]")},
			[][]int{{1}},
		},
		{
			"Case 3",
			args{tree.MustNewFromString("[]")},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
