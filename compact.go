package piscine

func Compact(ptr *[]string) int {
	count := 0
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			(*ptr)[count] = (*ptr)[i]
			count++
		}
	}
	(*ptr) = (*ptr)[:count]
	return count
}

//[ asdfasdf, asdfasdf, sadfasdf, , asdfasdf, sadf]
