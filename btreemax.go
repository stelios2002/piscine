package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	maxNode := root
	maxValue := maxNode.Data

	// Recursive function to traverse the tree and find the maximum value and its corresponding node
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Update maxNode and maxValue if the current node's value is greater
		if node.Data > maxValue {
			maxNode = node
			maxValue = node.Data
		}

		// Traverse left subtree
		traverse(node.Left)

		// Traverse right subtree
		traverse(node.Right)
	}

	// Start traversal from the root
	traverse(root)

	return maxNode
}
