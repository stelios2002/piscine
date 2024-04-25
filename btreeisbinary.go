package piscine

func isBST(node *TreeNode, min *string, max *string) bool {
	// Base case: an empty tree is a valid BST
	if node == nil {
		return true
	}

	// Check if the current node's value is within the valid range
	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}

	// Recursively check the left subtree with updated max value and right subtree with updated min value
	leftChild, rightChild := isBST(node.Left, min, &node.Data), isBST(node.Right, &node.Data, max)
	return leftChild && rightChild
}

func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}
