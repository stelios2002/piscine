package piscine

func LastRune(s string) rune {
	temp := []rune(s)
	return temp[len(temp)-1]
}
