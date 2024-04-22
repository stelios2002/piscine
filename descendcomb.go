package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a > 0; a-- {
		for b := a - 1; b >= 0; b-- {
			pnum(a, b)
			if a == 1 && b == 0 {
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func pnum(a, b int) {
	if a >= 10 {
		z01.PrintRune(rune(a/10) + '0')
		z01.PrintRune(rune(a%10) + '0')
	} else {
		z01.PrintRune('0')
		z01.PrintRune(rune(a) + '0')
	}
	z01.PrintRune(rune(' '))

	if b >= 10 {
		z01.PrintRune(rune(b/10) + '0')
		z01.PrintRune(rune(b%10) + '0')
	} else {
		z01.PrintRune('0')
		z01.PrintRune(rune(b) + '0')
	}
}
