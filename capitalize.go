package piscine

func Capitalize(s string) string {
	runeslice := []rune(s)
	inWord := false
	for i := 0; i < len(runeslice); i++ {
		if IsAlphaRune(string(runeslice[i])) {
			if inWord == false {
				runeslice[i] = ToUpperRune(runeslice[i]) // make it uppercase
				inWord = true
			} else {
				runeslice[i] = ToLowerRune(runeslice[i]) // make it lowercase
			}
		} else {
			inWord = false
		}
	}
	return string(runeslice)
}

func ToLowerRune(r rune) rune {
	if int(r) >= 65 && int(r) <= 90 {
		r = rune(int(r) + 32)
	}

	return r
}

func ToUpperRune(r rune) rune {
	if int(r) >= 97 && int(r) <= 122 {
		r = rune(int(r) - 32)
	}

	return r
}

func IsAlphaRune(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 97 && int(runeslice[i]) <= 122 || int(runeslice[i]) >= 65 && int(runeslice[i]) <= 90 || int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 {
		} else {
			return false
		}
	}
	return true
}
