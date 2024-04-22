package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runeslice := []rune(os.Args[0])
	// fmt.Println(str)
	for i := len(runeslice) - 1; i >= 0; i-- {
		if runeslice[i] == '/' {
			for x := i + 1; x < len(runeslice); x++ {
				z01.PrintRune(rune(runeslice[x]))
			}
			z01.PrintRune(rune('\n'))
			return
		}
	}
}
