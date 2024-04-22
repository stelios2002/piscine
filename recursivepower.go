package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	return Recur2(nb, power, nb)
}

func Recur2(nb, power, total int) int {
	power -= 1
	if power == 0 {
		return total
	} else {
		total *= nb
		return Recur2(nb, power, total)
	}
}
