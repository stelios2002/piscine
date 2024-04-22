package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	total := 1

	for i := 1; i <= nb; i++ {
		total *= i
		if total < 0 { // uf this ever becomes negative it means the value has overflowed and there's no point going forward. ('cause the task asks up not to)
			return 0
		}
	}
	return total
}
