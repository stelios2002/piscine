package piscine

func IsUpper(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 65 && int(runeslice[i]) <= 90 {
		} else {
			return false
		}
	}
	return true
}
