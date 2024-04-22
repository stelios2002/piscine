package piscine

func ListReverse(l *List) {
	myList := List{}
	if l.Head != nil {
		currnode := l.Head
		for currnode != l.Tail { // add all before tail
			ListPushFront(&myList, currnode.Data)
			currnode = currnode.Next
		}
		ListPushFront(&myList, currnode.Data) // Add tail
	}
	l.Head = myList.Head
	l.Tail = myList.Tail
}
