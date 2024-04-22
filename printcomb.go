package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := 0; a < 10; a++ {
		for b := a + 1; b < 10; b++ {
			for c := b + 1; c < 10; c++ {
				z01.PrintRune(rune(a + 48))
				z01.PrintRune(rune(b + 48))
				z01.PrintRune(rune(c + 48))
				if a == 7 && b == 8 && c == 9 {
					z01.PrintRune(rune('\n'))
					return
				}
				z01.PrintRune(rune(','))
				z01.PrintRune(rune(' '))
			}
		}
	}
}
