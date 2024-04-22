package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return l
	}
	curnode := l
	prevnode := l
	newnode := NodeI{}
	newnode.Data = data_ref
	if curnode.Data > newnode.Data {
		newnode.Next = curnode
		return &newnode
	}
	for curnode.Next != nil {
		prevnode = curnode
		curnode = curnode.Next
		if curnode.Data > newnode.Data {
			prevnode.Next = &newnode
			newnode.Next = curnode
			return l
		}
	}
	curnode.Next = &newnode
	return l
}
