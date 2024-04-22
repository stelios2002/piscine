package piscine

//    If the number is even, divide it by two.
// If the number is odd, triple it and add one.

func CollatzCountdown(start int) int {
	if start < 1 {
		return -1
	}
	counter := 0
	for start > 1 {
		if start%2 == 0 {
			start /= 2
		} else if start%2 == 1 {
			start = 3*start + 1
		}
		counter++
		if start == 1 {
			break
		}
	}
	return counter
}
