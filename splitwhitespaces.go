package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	starti := 0
	endi := 0
	foundWhiteSpace := true
	for i := 0; i < len(s); i++ {
		if isWhitespaceChar(s[i]) == true {
			if foundWhiteSpace == true {
				//
			} else {
				// add word if word was being recorded
				endi = i
				result = append(result, s[starti:endi])
				starti = i
			}
			foundWhiteSpace = true

		} else {
			if foundWhiteSpace {
				foundWhiteSpace = false
				// start recording word until whitespace found
				starti = i
				endi = i + 1
			}
		}
	}

	if endi > starti {
		result = append(result, s[starti:])
	}

	return result
}

var whiteSpaceChars = []byte{' ', '	', '\n'}

func isWhitespaceChar(char byte) bool {
	for i := 0; i < 3; i++ {
		if whiteSpaceChars[i] == char {
			return true
		}
	}
	return false
}
