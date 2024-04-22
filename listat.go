package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos == 0 {
		return l
	} else if pos > 0 {
		size := 0
		for l.Next != nil {
			size++
			l = l.Next
			if size == pos {
				return l
			}
		}
	}
	return nil
}
