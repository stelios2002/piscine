package piscine

func ListLast(l *List) interface{} {
	size := 0
	if l.Head != nil {
		size++
		curnode := l.Head
		for curnode != l.Tail {
			curnode = curnode.Next
			size++
		}
	}

	if size == 0 {
		return nil
	} else if size == 1 {
		if l.Head != nil {
			return l.Head.Data
		} else if l.Tail != nil {
			return l.Tail.Data
		}
	}
	return l.Tail.Data
}
