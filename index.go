package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 || len(toFind) > len(s) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			found := true
			for subi := 0; subi < len(toFind); subi++ {
				if subi+i >= len(s) {
					return -1
				}
				if s[subi+i] != toFind[subi] {
					found = false
					break
				}
			}
			if found == true {
				return i
			}
		}
	}
	return -1
}
