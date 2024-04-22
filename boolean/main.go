package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

var (
	EvenMsg string = "I have an even number of arguments"
	OddMsg  string = "I have an odd number of arguments"
)

func main() {
	lengthOfArg := len(os.Args) - 1

	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
