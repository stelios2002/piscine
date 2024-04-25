package piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	parent := root
	child := TreeNode{}
	child.Data = data
	for {
		if data > parent.Data {
			// go right
			if parent.Right != nil {
				parent = parent.Right
				continue
			} else {
				child.Parent = parent
				parent.Right = &child
				// insert it on the right
				break
			}
		} else {
			// go left
			if parent.Left != nil {
				parent = parent.Left
				continue
			} else {
				child.Parent = parent
				parent.Left = &child
				// insert it on the right
				break
			}
		}
	}
	return root
}
