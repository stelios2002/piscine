package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		return rplc
	}
	parent := BTreeSearchParent(root, node)
	if parent == nil {
		return root
	}
	rplc.Parent = node.Parent
	if parent.Left == node {
		parent.Left = rplc
	} else {
		parent.Right = rplc
	}

	return root
}

func BTreeSearchParent(root, node *TreeNode) *TreeNode {
	if root == nil || node == root {
		return nil
	}

	if (root.Left == node) || (root.Right == node) {
		return root
	}

	if root.Data > node.Data {
		return BTreeSearchParent(root.Left, node)
	}

	return BTreeSearchParent(root.Right, node)
}
