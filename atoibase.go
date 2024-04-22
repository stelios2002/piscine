package piscine

var mymap = make(map[rune]int)

func AtoiBase(s string, base string) int {
	basesize := len(base)

	if len(base) < 2 {
		return 0
	}

	for i := 0; i < len(base); i++ {
		if rune(base[i]) == rune('-') || rune(base[i]) == rune('+') {
			return 0
		}
		count := 0
		for i2 := 0; i2 < len(base); i2++ {
			if base[i] == base[i2] {
				count++ // non unqiue characters not working
			}
		}
		if count >= 2 {
			return 0
		}
	}

	for i := 0; i < basesize; i++ {
		mymap[rune(base[i])] = i
	}

	total := 0

	for i := 0; i < len(s); i++ {
		power := 1

		for powi := 1; powi < len(s)-i; powi++ {
			power *= basesize
		}
		total += power * mymap[rune(s[i])]

	}

	return total
}
