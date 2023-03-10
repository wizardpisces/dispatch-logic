package prune_tree

import (
	"github.com/wizardpisces/dispatch-logic/structures"
)

func PruneTree[T int](root *structures.TreeNode[T]) *structures.TreeNode[T] {
	if root == nil || root.Val == nil {
		return nil
	}
	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)
	if root.Left == nil && root.Right == nil && *root.Val == 0 {
		return nil
	}
	return root
}
