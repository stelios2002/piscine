package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}

	setPoint(&points)

	a := "x = "
	b := ", y = "

	c := '\n'

	for _, c := range a {
		z01.PrintRune(c)
	}

	PrintNbr(points.x)

	for _, c := range b {
		z01.PrintRune(c)
	}

	PrintNbr(points.y)

	z01.PrintRune(c)
}

func PrintNbr(n int) {
	temp := n

	if n == -9223372036854775808 {
		temp = 8
		n += 1
		n = -n
		z01.PrintRune(rune('-'))

	} else {

		if n < 0 {
			z01.PrintRune(rune('-'))
			n = -n
		}
		temp = n % 10
	}

	n /= 10

	if n > 0 {
		PrintNbr(n)
	}

	z01.PrintRune(rune(temp + 48))

	return
}
