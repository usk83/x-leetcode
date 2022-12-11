package tree

import (
	"reflect"
	"testing"
)

func intAddr(n int) *int {
	return &n
}

func TestNew(t *testing.T) {
	type args struct {
		arr []*int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"Simple",
			args{[]*int{intAddr(1), intAddr(2), intAddr(3), intAddr(4), intAddr(5), intAddr(6), intAddr(7)}},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			"Normal tree 01",
			args{[]*int{intAddr(3), intAddr(9), intAddr(20), nil, nil, intAddr(15), intAddr(7)}},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			"Normal tree 02",
			args{[]*int{intAddr(1), nil, intAddr(2), nil, intAddr(3)}},
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			"Normal tree 03",
			args{[]*int{intAddr(1), nil, intAddr(2), nil, intAddr(3), intAddr(1), intAddr(2), nil, nil, intAddr(1)}},
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
				},
			},
		},
		{
			"Empty array",
			args{[]*int{}},
			nil,
		},
		{
			"Nil input",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *TreeNode
		wantErr bool
	}{
		{
			"Simple",
			args{"[1,2,3,4,5,6,7]"},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
		{
			"Normal tree 01",
			args{"[3,9,20,null,null,15,7]"},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
		{
			"Empty array",
			args{"[]"},
			nil,
			false,
		},
		{
			"Empty input",
			args{""},
			nil,
			false,
		},
		{
			"Invalid string format",
			args{"[1,2,3}"},
			nil,
			true,
		},
		{
			"Contains invalid value",
			args{"[523,nil,5]"},
			nil,
			true,
		},
		{
			"Contains empty value",
			args{"[124,5231,,-421]"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFromString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_String(t *testing.T) {
	tests := []struct {
		name     string
		receiver *TreeNode
		want     string
	}{
		{
			"Simple",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			"[1,2,3,4,5,6,7]",
		},
		{
			"Normal tree 01",
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			"[3,9,20,null,null,15,7]",
		},
		{
			"Nil node on the end",
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			"[1,2]",
		},
		{
			"Nil receiver",
			nil,
			"[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.receiver.String(); got != tt.want {
				t.Errorf("ListNode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
