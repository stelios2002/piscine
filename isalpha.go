package piscine

func IsAlpha(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 97 && int(runeslice[i]) <= 122 || int(runeslice[i]) >= 65 && int(runeslice[i]) <= 90 || int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 {
		} else {
			return false
		}
	}
	return true
}
