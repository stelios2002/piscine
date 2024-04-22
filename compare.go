package piscine

func Compare(a, b string) int {
	size := len(b)
	if len(a) < len(b) {
		size = len(a)
	}

	for i := 0; i < size; i++ {
		if a[0] > b[0] {
			return 1
		} else if a[0] < b[0] {
			return -1
		}
	}
	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	}
	return 0
}
