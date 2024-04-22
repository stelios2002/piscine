package piscine

func ListPushFront(l *List, data interface{}) {
	size := 0
	if l.Head != nil {
		size++
		curnode := l.Head
		for curnode != l.Tail {
			curnode = curnode.Next
			size++
		}
	}
	newnode := NodeL{}
	newnode.Data = data
	if size == 0 {
		l.Head = &newnode
		l.Tail = &newnode
	} else if size == 1 {
		newnode.Next = l.Head
		l.Head = &newnode
	} else {
		newnode.Next = l.Head
		l.Head = &newnode
	}
}
