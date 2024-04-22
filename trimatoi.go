package piscine

func TrimAtoi(s string) int {
	runeslice := []rune(s)
	newruneslice := []rune{}
	hasNumBeenFound := false
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 || int(runeslice[i]) == int(rune('+')) || int(runeslice[i]) == int(rune('-')) { // also add case for - or +

			if int(runeslice[i]) == int(rune('+')) || int(runeslice[i]) == int(rune('-')) {
				if hasNumBeenFound == false {
					newruneslice = append(newruneslice, runeslice[i])
				}
			} else {
				newruneslice = append(newruneslice, runeslice[i])
			}

			if int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 {
				hasNumBeenFound = true
			}

		}
	}

	return Atoi2(string(newruneslice))
}

var valids3 string

func Atoi2(s string) int {
	if len(s) == 0 {
		return 0
	}
	symbol_count := 0
	valids3 = "0123456789-+"
	for si := 0; si < len(s); si++ {
		FoundMatch := false
		for vi := 0; vi < len(valids3); vi++ {
			if valids3[vi] == s[si] {
				if s[si] == valids3[10] || s[si] == valids3[11] {
					symbol_count++
					if symbol_count > 1 || si > 0 {
						return 0
					}
				}

				FoundMatch = true
				break
			}
		}
		if FoundMatch == false {
			return 0
		}
	}
	is_neg := false
	if s[0] == valids3[10] {
		is_neg = true
	}

	temp := 0
	sig := 1

	for i := len(s) - 1; i >= symbol_count; i-- {
		temp += sig * (int(rune(s[i]) - 48))
		sig *= 10
	}

	if is_neg == true {
		temp = -temp
	}

	return temp
}
