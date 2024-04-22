package piscine

import (
	"github.com/01-edu/z01"
)

var (
	mySlice  []int
	mySlice2 []int
)

var once bool

func PrintCombN(n int) {
	once = true
	for i := 0; i < n; i++ {
		mySlice = append(mySlice, 0)
		mySlice2 = append(mySlice2, 0)
	}

	deeper(n, -1)

	for i := 0; i < len(mySlice); i++ {
		z01.PrintRune(rune(mySlice2[i] + 48))
	}
	z01.PrintRune('\n')
	mySlice = mySlice[:0]
	mySlice2 = mySlice2[:0]
	return
}

func deeper(maxd int, num int) {
	maxd -= 1
	for i := num + 1; i < 10; i++ {
		mySlice[len(mySlice)-1-maxd] = i
		if maxd > 0 {
			deeper(maxd, i)
		} else {
			print_nums()
		}
	}
}

func print_nums() {
	if once == true {
		once = false
	} else {
		for i := 0; i < len(mySlice); i++ {
			z01.PrintRune(rune(mySlice2[i] + 48))
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	for i := 0; i < len(mySlice); i++ {
		mySlice2[i] = mySlice[i]
	}
}
