package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	newdir := 0
	dir := 0

	for i := 0; i < len(a)-1; i++ {

		newres := f(a[i], a[i+1])
		if newres != 0 {

			if newres > 0 {
				newdir = 1
			} else {
				newdir = -1
			}

			if dir == 0 {
				dir = newdir
			} else {
				if dir != newdir {
					return false
				}
			}
		}
	}

	return true
}
