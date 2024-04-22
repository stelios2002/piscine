package piscine

func AlphaCount(s string) int {
	// >= 65 && <= 122
	counter := 0
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 65 && int(runeslice[i]) <= 90 || int(runeslice[i]) >= 97 && int(runeslice[i]) <= 122 {
			counter++
		}
	}
	return counter
}
