package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := 1; i < len(args); i++ {
		for ch_i := 0; ch_i < len(args[i]); ch_i++ {
			z01.PrintRune(rune(args[i][ch_i]))
		}
		z01.PrintRune(rune('\n'))
	}
}
