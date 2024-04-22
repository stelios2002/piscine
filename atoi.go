package piscine

// import "fmt"

var valids2 string

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	symbol_count := 0
	valids2 = "0123456789-+"
	for si := 0; si < len(s); si++ {
		FoundMatch := false
		for vi := 0; vi < len(valids2); vi++ {
			if valids2[vi] == s[si] {
				if s[si] == valids2[10] || s[si] == valids2[11] {
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
	if s[0] == valids2[10] {
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
