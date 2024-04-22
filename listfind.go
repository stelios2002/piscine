package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head != nil {
		curnode := l.Head
		if comp(curnode.Data, ref) {
			return &curnode.Data
		}
		for curnode != l.Tail {
			curnode = curnode.Next
			if comp(curnode.Data, ref) {
				return &curnode.Data
			}
		}
	}
	return nil
}
