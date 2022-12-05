package list

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(arr []int) *ListNode {
	metaNode := &ListNode{}

	curNode := metaNode
	for _, a := range arr {
		curNode.Next = &ListNode{Val: a}
		curNode = curNode.Next
	}

	return metaNode.Next
}

func NewFromString(str string) (*ListNode, error) {
	if len(str) == 0 {
		return nil, nil
	}

	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		return nil, &InvalidStringError{}
	}

	str = str[1 : len(str)-1]

	if len(str) == 0 {
		return nil, nil
	}

	vs := strings.Split(str, ",")

	arr := make([]int, len(vs))
	for i, v := range vs {
		a, err := strconv.Atoi(v)
		if err != nil {
			return nil, &InvalidStringError{err}
		}

		arr[i] = a
	}

	return New(arr), nil
}

func MustNewFromString(str string) *ListNode {
	n, err := NewFromString(str)
	if err != nil {
		panic(err)
	}

	return n
}

func (n *ListNode) String() string {
	var vs []string
	for node := n; node != nil; node = node.Next {
		vs = append(vs, strconv.Itoa(node.Val))
	}

	return fmt.Sprintf("[%s]", strings.Join(vs, ","))
}
