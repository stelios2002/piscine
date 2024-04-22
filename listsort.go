package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	size := 0
	if l == nil {
		return l
	}
	if l.Next != nil {
		size++
		curnode := l
		for curnode.Next != nil {
			curnode = curnode.Next
			size++
		}
	}
	if size < 2 {
		return l
	}
	head := l
	for i := 0; i < size; i++ {
		curr := l
		for x := i + 1; x < size; x++ {
			if curr.Data > curr.Next.Data {
				curr.Data, curr.Next.Data = curr.Next.Data, curr.Data
			}
			curr = curr.Next
		}
		l = head
	}
	return head
}
