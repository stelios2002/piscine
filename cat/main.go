package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			dat, err := os.ReadFile(arg)

			if err != nil {
				// Handle the error if the file cannot be opened
				// fmt.Println("Error:", err)
				PrintStr("ERROR: open " + arg + ": no such file or directory\n")
				os.Exit(1)
				return
			} else {
				PrintStr(string(dat))
			}
		}
	} else {

		input := make([]byte, 0)
		buffer := make([]byte, 1024)
		// Read from os.Stdin into buffer
		n, _ := os.Stdin.Read(buffer)

		input = append(input, buffer[:n]...)
		PrintStr(string(input))

	}
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
