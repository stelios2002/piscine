package main

import (
	"os"

	"github.com/01-edu/z01"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func main() {
	fullWord := ""
	if len(os.Args) > 1 {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if i+1 < len(args) {
				fullWord = fullWord + args[i] + " "
			} else {
				fullWord = fullWord + args[i]
			}
		}
		vowelSlice := []rune{}
		fullWordSlice := []rune(fullWord)
		for i := 0; i < len(fullWord); i++ {
			if isVowel(rune(fullWord[i])) {
				vowelSlice = append(vowelSlice, rune(fullWord[i]))
			}
		}
		vCounter := 0
		for i := len(fullWordSlice) - 1; i >= 0; i-- {
			if isVowel(rune(fullWord[i])) {
				fullWordSlice[i] = rune(vowelSlice[vCounter])
				vCounter++
			}
		}

		for i := 0; i < len(fullWordSlice); i++ {
			z01.PrintRune(rune(fullWordSlice[i]))
		}

		z01.PrintRune('\n')

	} else {
		z01.PrintRune('\n')
	}
}

func isVowel(r rune) bool {
	for i := 0; i < len(vowels); i++ {
		if vowels[i] == r {
			return true
		}
	}
	return false
}
