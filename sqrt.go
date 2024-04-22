package piscine

func Sqrt(nb int) int { // using the subtraction method
	subtractor := 1
	counter := 0
	for nb > 0 {
		nb -= subtractor
		subtractor += 2
		counter++
	}
	if nb == 0 {
		return counter
	}
	return 0
}
