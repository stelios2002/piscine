package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	newword := ""
	counter := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			counter++
		} else {
			break
		}
	}
	for i := 0; i < len(str); i++ {
		ch := string(str[i])
		if ch == string(0) {
		}
		if str[i] != ' ' {
			counter++
			if (counter)%6 != 0 {
				newword = newword + string(str[i])
			} else {
				newword = newword + " "
			}
		} else {
			counter++
			if (counter)%6 == 0 {
				newword = newword + " "
			} else {
				counter--
			}
			// for str[i] == ' ' {
			// 	i++
			// }
		}
	}
	return newword + "\n"
}
