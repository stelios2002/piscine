package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		if root.Right != nil {
			BTreeApplyInorder(root.Right, f)
		}
	} else if root.Right != nil {
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	} else {
		f(root.Data)
	}
}
