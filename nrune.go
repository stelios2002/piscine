package piscine

func NRune(s string, n int) rune {
	temp := []rune(s)
	if n <= 0 || n > len(temp) {
		return 0
	}
	return temp[n-1]
}
