package piscine

func IsNumeric(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 || int(runeslice[i]) == '-' || int(runeslice[i]) == '+' {
		} else {
			return false
		}
	}
	return true
}
