package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := AtoiBase2Dec(nbr, baseFrom)
	res2 := Dec2ArbitraryBase(res, baseTo)
	return res2
}

func Dec2ArbitraryBase(nbr int, base string) string {
	if len(base) < 2 {
		return ""
	}

	for i := 0; i < len(base); i++ {
		if rune(base[i]) == rune('-') || rune(base[i]) == rune('+') {
			return ""
		}
		count := 0
		for i2 := 0; i2 < len(base); i2++ {
			if base[i] == base[i2] {
				count++ // non unqiue characters not working
			}
		}
		if count >= 2 {
			return ""
		}
	}

	store := 0
	isneg := false

	if nbr == 0 {
		// z01.PrintRune(rune(base[0]))
	}
	if nbr < 0 {
		isneg = true
		nbr = -nbr
	}
	if isneg == true {
		// z01.PrintRune(rune('-'))
	}

	result := ""
	for nbr > 0 || nbr < 0 {
		result = string(base[abs2(nbr%len(base))]) + result
		nbr /= len(base)
		if store > 0 {
			nbr += 1
		}
	}
	return result
}

func abs2(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

var mymap2 = make(map[rune]int)

func AtoiBase2Dec(s string, base string) int {
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
		mymap2[rune(base[i])] = i
	}

	total := 0

	for i := 0; i < len(s); i++ {
		power := 1

		for powi := 1; powi < len(s)-i; powi++ {
			power *= basesize
		}
		total += power * mymap2[rune(s[i])]

	}

	return total
}
