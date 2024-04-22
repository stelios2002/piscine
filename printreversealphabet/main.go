package main

import (
	"github.com/01-edu/z01"
)

func main() {
	myString := "abcdefghijklmnopqrstuvwxyz"
	for i := (len(myString) - 1); i >= 0; i-- {
		z01.PrintRune(rune(myString[i]))
	}
	z01.PrintRune(rune('\n'))

	// fmt.Println("Hello, World!")
}
