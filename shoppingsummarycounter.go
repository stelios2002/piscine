package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	mymap := map[string]int{}

	if str == "" {
		mymap[""] = 1
		return mymap
	} else if str == " " {
		mymap[""] = 2
		return mymap
	}

	if len(str) >= 1 {
		if string(str[0]) == " " {
			mymap[""]++
		}
		if string(str[len(str)-1]) == " " {
			mymap[""]++
		}
	}

	sslice := Split3(str, " ")

	for _, word := range sslice {
		mymap[word]++
	}

	return mymap
}

func Split3(s, sep string) []string {
	result := []string{}
	for len(s) > 0 {
		end := Index3(s, sep)
		if end != -1 {
			result = append(result, s[0:end])
			s = s[len(sep)+end:]
		} else {
			result = append(result, s[0:])
			s = ""
		}
	}

	return result
}

func Index3(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 || len(toFind) > len(s) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			found := true
			for subi := 0; subi < len(toFind); subi++ {
				if subi+i >= len(s) {
					return -1
				}
				if s[subi+i] != toFind[subi] {
					found = false
					break
				}
			}
			if found == true {
				return i
			}
		}
	}
	return -1
}
