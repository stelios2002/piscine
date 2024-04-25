package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		if root.Right != nil {
			BTreeApplyPreorder(root.Right, f)
		}
	} else if root.Right != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Right, f)
	} else {
		f(root.Data)
	}
}
