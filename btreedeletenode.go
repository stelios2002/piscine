package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Case 1: Node to be deleted has no children
	if node.Left == nil && node.Right == nil {
		return deleteLeaf(root, node)
	}

	// Case 2: Node to be deleted has one child
	if node.Left == nil {
		return promoteChild(root, node, node.Right)
	}
	if node.Right == nil {
		return promoteChild(root, node, node.Left)
	}

	// Case 3: Node to be deleted has two children
	successor := BTreeMinBST(node.Right)
	node.Data = successor.Data
	node.Right = BTreeDeleteNode(node.Right, successor)
	return root
}

func deleteLeaf(root, node *TreeNode) *TreeNode {
	if root == node {
		return nil
	}

	parent := BTreeSearchParent(root, node)
	if parent == nil {
		return root
	}

	if parent.Left == node {
		parent.Left = nil
	} else {
		parent.Right = nil
	}

	return root
}

func promoteChild(root, node, child *TreeNode) *TreeNode {
	if root == node {
		return child
	}

	parent := BTreeSearchParent(root, node)
	if parent == nil {
		return root
	}

	if parent.Left == node {
		parent.Left = child
	} else {
		parent.Right = child
	}

	return root
}

func BTreeMinBST(root *TreeNode) *TreeNode {
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
