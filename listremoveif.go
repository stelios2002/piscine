package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	// trim the head until a non match comes across
	for l.Head.Data == data_ref {
		if l.Head.Next != nil {
			l.Head = l.Head.Next
		} else {
			l.Head = nil
			return
		}
	}
	if l.Head != nil {
		curnode := l.Head
		// previous needs to be stored in order to successfully remove the node from the middle of the linked list
		prevnode := curnode
		for curnode != l.Tail {
			prevnode = curnode
			curnode = curnode.Next
			if curnode == l.Tail {
				break
			}
			if curnode.Data == data_ref {
				prevnode.Next = curnode.Next
				curnode = prevnode
			}
		}
		if l.Tail.Data == data_ref {
			l.Tail = prevnode
			l.Tail.Next = nil
		}
	}
}
