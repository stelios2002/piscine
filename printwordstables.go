package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		for _, r := range str {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
