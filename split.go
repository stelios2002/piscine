package piscine

func Split(s, sep string) []string {
	result := []string{}
	for len(s) > 0 {
		end := Index2(s, sep)
		if end != -1 {
			result = append(result, s[0:end])
			s = s[len(sep)+end:]
		} else {
			result = append(result, s[0:])
			s = ""
		}
	}

	return result
}

func Index2(s string, toFind string) int {
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
