package list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"Simple",
			args{[]int{0, 1, 2}},
			&ListNode{0, &ListNode{1, &ListNode{2, nil}}},
		},
		{
			"Contains negative value",
			args{[]int{-10, 777, 0, 1}},
			&ListNode{-10, &ListNode{777, &ListNode{0, &ListNode{1, nil}}}},
		},
		{
			"Empty array",
			args{[]int{}},
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
		want    *ListNode
		wantErr bool
	}{
		{
			"Normal",
			args{"[124,5231,0312,-421]"},
			&ListNode{124, &ListNode{5231, &ListNode{312, &ListNode{-421, nil}}}},
			false,
		},
		{
			"Single value",
			args{"[111]"},
			&ListNode{111, nil},
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
			args{"[523,abc,5]"},
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

func TestNode_String(t *testing.T) {
	tests := []struct {
		name     string
		receiver *ListNode
		want     string
	}{
		{
			"Simple",
			&ListNode{134, &ListNode{9, &ListNode{33, nil}}},
			"[134,9,33]",
		},
		{
			"Contains negative value",
			&ListNode{-10, &ListNode{777, &ListNode{0, &ListNode{1, nil}}}},
			"[-10,777,0,1]",
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
