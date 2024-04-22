package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 0; i < 99; i++ {
		for x := i + 1; x < 100; x++ {
			if i >= 10 {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))
			} else {
				z01.PrintRune(48)
				z01.PrintRune(rune(i + 48))
			}

			z01.PrintRune(' ')

			if x >= 10 {
				z01.PrintRune(rune(x/10 + 48))
				z01.PrintRune(rune(x%10 + 48))
			} else {
				z01.PrintRune(48)
				z01.PrintRune(rune(x + 48))
			}
			if i == 98 && x == 99 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
