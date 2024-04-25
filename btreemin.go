package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	minNode := root
	minValue := minNode.Data

	// Recursive function to traverse the tree and find the maximum value and its corresponding node
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Update minNode and minValue if the current node's value is greater
		if node.Data < minValue {
			minNode = node
			minValue = node.Data
		}

		// Traverse left subtree
		traverse(node.Left)

		// Traverse right subtree
		traverse(node.Right)
	}

	// Start traversal from the root
	traverse(root)

	return minNode
}
