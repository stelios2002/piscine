package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	total := 1
	total = Recur(total, nb, 1)

	return total
}

func Recur(total, nb, nb_counter int) int {
	total *= nb_counter
	if total < 0 {
		return 0
	}
	nb_counter += 1
	if nb_counter > nb {
		return total
	} else {
		total = Recur(total, nb, nb_counter)
		return total
	}
}
