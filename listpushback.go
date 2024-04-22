package piscine

func ListPushBack(l *List, data interface{}) {
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
		l.Head.Next = &newnode
		l.Tail = &newnode
	} else {
		l.Tail.Next = &newnode
		l.Tail = &newnode
	}
}
