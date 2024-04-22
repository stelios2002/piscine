package piscine

func BasicJoin(elems []string) string {
	newstring := ""
	for i := 0; i < len(elems); i++ {
		newstring = newstring + elems[i]
	}
	return newstring
}
