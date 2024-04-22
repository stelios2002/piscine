package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	// args = []string{"1", "2", "3"}
	if len(args) > 1 {
		for i := len(args) - 1; i > 0; i-- {
			for ch_i := 0; ch_i < len(args[i]); ch_i++ {
				z01.PrintRune(rune(args[i][ch_i]))
			}
			z01.PrintRune(rune('\n'))
		}
	}
}
