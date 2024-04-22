package piscine

func DescendAppendRange(max, min int) []int {
	slice := []int{}
	if max <= min {
		return slice
	}
	for i := max; i > min; i-- {
		slice = append(slice, i)
	}
	return slice
}
