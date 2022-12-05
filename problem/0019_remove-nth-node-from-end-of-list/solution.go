package remove_nth_node_from_end_of_list

import "github.com/usk83/x-leetcode/internal/ds/list"

// Constraints:
//   - The number of nodes in the list is sz.
//   - 1 <= sz <= 30
//   - 0 <= Node.val <= 100
//   - 1 <= n <= sz

var (
	_                = twoPass
	removeNthFromEnd = onePass
)

// count the number of nodes in the list and determine the node to remove
//   - Time Complexity: O(sz)
//   - Space Complexity: O(1)
func twoPass(head *list.ListNode, n int) *list.ListNode {
	var sz int
	for node := head; node != nil; node = node.Next {
		sz++
	}

	metaNode := &list.ListNode{Val: -1, Next: head}
	curNode := metaNode
	for i := n; i < sz; i++ {
		curNode = curNode.Next
	}

	curNode.Next = curNode.Next.Next

	return metaNode.Next
}

// keep tracking of the node to the left of the node to remove
//   - Time Complexity: O(sz)
//   - Space Complexity: O(1)
func onePass(head *list.ListNode, n int) *list.ListNode {
	metaNode := &list.ListNode{Val: -1, Next: head}

	pointer, counter := metaNode, -1
	for node := metaNode; node != nil; node = node.Next {
		if counter++; counter > n {
			pointer = pointer.Next
		}
	}

	pointer.Next = pointer.Next.Next

	return metaNode.Next
}
