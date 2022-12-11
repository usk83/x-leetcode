package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type InvalidStringError struct {
	detail error
}

func (e *InvalidStringError) Error() string {
	const msg = "cannot parse input string"

	if e.detail != nil {
		return fmt.Sprintf("%s: %v", msg, e.detail)
	}

	return msg
}

func (e *InvalidStringError) Unwrap() error {
	return e.detail
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}

	q, left := []*TreeNode{root}, true
	for _, v := range arr[1:] {
		if v != nil {
			n := &TreeNode{Val: *v}
			if left {
				q[0].Left = n
			} else {
				q[0].Right = n
			}
			q = append(q, n)
		}

		if !left {
			q = q[1:]
		}

		left = !left

		if len(q) == 0 {
			break
		}
	}

	return root
}

func NewFromString(str string) (*TreeNode, error) {
	if len(str) == 0 || str == "[]" {
		return nil, nil
	}

	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		return nil, &InvalidStringError{}
	}

	vs := strings.Split(str[1:len(str)-1], ",")

	arr := make([]*int, len(vs))
	for i, v := range vs {
		if v == "null" {
			continue
		}

		a, err := strconv.Atoi(v)
		if err != nil {
			return nil, &InvalidStringError{err}
		}

		arr[i] = &a
	}

	return New(arr), nil
}

func MustNewFromString(str string) *TreeNode {
	n, err := NewFromString(str)
	if err != nil {
		panic(err)
	}

	return n
}

func (n *TreeNode) String() string {
	if n == nil {
		return "[]"
	}

	var (
		vs                   = []string{strconv.Itoa(n.Val)}
		q                    = []*TreeNode{n}
		lastNonLeafNodeIndex = 0 // last index of a non-leaf node
	)
	for len(q) != 0 && lastNonLeafNodeIndex >= 0 {
		for i, tn := range []*TreeNode{q[0].Left, q[0].Right} {
			if tn == nil {
				if !(lastNonLeafNodeIndex == 0 && i == 1) {
					vs = append(vs, "null")
				}
			} else {
				vs = append(vs, strconv.Itoa(tn.Val))
				q = append(q, tn)
				if tn.Left != nil || tn.Right != nil {
					lastNonLeafNodeIndex = len(q) - 1
				}
			}
		}
		q = q[1:]
		lastNonLeafNodeIndex--
	}

	return fmt.Sprintf("[%s]", strings.Join(vs, ","))
}
