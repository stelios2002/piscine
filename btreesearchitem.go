package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root.Data == elem {
		return root
	}
	if root.Left != nil {
		ret := BTreeSearchItem(root.Left, elem)
		if ret != nil {
			return ret
		}
		if root.Right != nil {
			ret = BTreeSearchItem(root.Right, elem)
			if ret != nil {
				return ret
			}
		}
	} else if root.Right != nil {
		ret := BTreeSearchItem(root.Right, elem)
		if ret != nil {
			return ret
		}
	}
	return nil
}
