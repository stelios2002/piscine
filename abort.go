package piscine

func Abort(a, b, c, d, e int) int {
	s := []int{a, b, c, d, e}

	for ai := 0; ai < len(s)-1; ai++ {
		for bi := ai; bi < len(s); bi++ {
			if s[ai] > s[bi] {
				s[ai], s[bi] = s[bi], s[ai]
			}
		}
	}
	return s[2]
}
