package piscine

func Unmatch(a []int) int {
	mymap := map[int]int{}
	ret := -1
	for _, i := range a {
		_, ok := mymap[i]
		if ok {
			mymap[i] += 1
		} else {
			mymap[i] = 1
		}

	}

	for _, i := range a {
		val, _ := mymap[i]
		if val%2 == 1 {
			ret = i
			break
		}

	}
	return ret
}
