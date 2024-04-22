package piscine

var valids string

func BasicAtoi2(s string) int {
	valids = "0123456789-"
	for si := 0; si < len(s); si++ {
		FoundMatch := false
		for vi := 0; vi < len(valids); vi++ {
			if valids[vi] == s[si] {
				FoundMatch = true
				break
			}
		}
		if FoundMatch == false {
			return 0
		}
	}

	temp := 0
	sig := 1
	for i := len(s) - 1; i >= 0; i-- {
		temp += sig * (int(rune(s[i]) - 48))
		sig *= 10
	}
	return temp
}
