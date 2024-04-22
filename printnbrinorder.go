package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	} else if n == 0 {
		z01.PrintRune(rune(48))
	}

	arr := [10]int{}
	for n > 0 {
		arr[n%10] += 1
		n /= 10
	}

	for i := 0; i < 10; i++ {
		for x := 0; x < arr[i]; x++ {
			z01.PrintRune(rune(48 + i))
		}
	}
}
