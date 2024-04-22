package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	result := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		result[i] = i + min
	}

	return result
}
