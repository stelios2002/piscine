package piscine

func ListSize(l *List) int {
	size := 0
	if l.Head != nil {
		size++
		curnode := l.Head
		for curnode != l.Tail {
			curnode = curnode.Next
			size++
		}
	}
	return size
}
