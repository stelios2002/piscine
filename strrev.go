package piscine

func StrRev(s string) string {
	RuneSlice := []rune(s)
	r := len(RuneSlice) - 1

	for l := 0; l < len(RuneSlice)/2; l++ {
		Swap2(&RuneSlice[l], &RuneSlice[r])
		r -= 1
	}

	return string(RuneSlice)
}

func Swap2(a *rune, b *rune) { // LMAO EVEN
	temp := *b
	*b = *a
	*a = temp
}
