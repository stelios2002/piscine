package piscine

func ListMerge(l1 *List, l2 *List) {
	if !listIsValid(l1) && listIsValid(l2) {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else if listIsValid(l1) && listIsValid(l2) {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}

func listIsValid(l *List) bool {
	if l.Head != nil && l.Tail != nil {
		return true
	}
	return false
}
