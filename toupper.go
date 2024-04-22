package piscine

func ToUpper(s string) string {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 97 && int(runeslice[i]) <= 122 {
			runeslice[i] = rune(int(runeslice[i]) - 32)
		}
	}
	return string(runeslice)
}
