package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left) + 1
	rightHeight := BTreeLevelCount(root.Right) + 1
	if leftHeight < rightHeight {
		return rightHeight
	}
	return leftHeight
}
