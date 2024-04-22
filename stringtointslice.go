package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return []int(nil)
	}
	slice := []int{}
	runeslice := []rune(str)
	for i := 0; i < len(runeslice); i++ {
		slice = append(slice, int(runeslice[i]))
	}
	return slice
}
