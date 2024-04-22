package piscine

// stole solution from another's chatgpt solution :/

func JumpOver(str string) string {
	s := ""
	for i, ch := range str {
		if (i+1)%3 == 0 {
			s = s + string(ch)
		}
	}
	if s == "" {
		return string('\n')
	}
	s = s + "\n"
	return s
}
