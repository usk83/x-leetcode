package binary_tree_level_order_traversal

import "github.com/usk83/x-leetcode/internal/ds/tree"

// Constraints:
//   - The number of nodes in the tree is in the range [0, 2000].
//   - -1000 <= Node.val <= 1000

func levelOrder(root *tree.TreeNode) [][]int {
	leveledVals := [][]int{}

	if root == nil {
		return leveledVals
	}

	currentLevelNodes := []*tree.TreeNode{root}
	for len(currentLevelNodes) != 0 {
		var (
			vals           = make([]int, len(currentLevelNodes))
			nextLevelNodes []*tree.TreeNode
		)
		for i, n := range currentLevelNodes {
			vals[i] = int(n.Val)

			if n.Left != nil {
				nextLevelNodes = append(nextLevelNodes, n.Left)
			}

			if n.Right != nil {
				nextLevelNodes = append(nextLevelNodes, n.Right)
			}
		}
		leveledVals = append(leveledVals, vals)
		currentLevelNodes = nextLevelNodes
	}

	return leveledVals
}
