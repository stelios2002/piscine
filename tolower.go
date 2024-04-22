package piscine

func ToLower(s string) string {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 65 && int(runeslice[i]) <= 90 {
			runeslice[i] = rune(int(runeslice[i]) + 32)
		}
	}
	return string(runeslice)
}
