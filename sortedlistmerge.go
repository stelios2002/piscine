package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	for {
		n1 = SortListInsert(n1, n2.Data)
		n2 = n2.Next
		if n2.Next == nil {
			n1 = SortListInsert(n1, n2.Data)
			break
		}
	}
	return n1
}
