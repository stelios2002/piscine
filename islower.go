package piscine

func IsLower(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 97 && int(runeslice[i]) <= 122 {
		} else {
			return false
		}
	}
	return true
}
