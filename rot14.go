package piscine

func Rot14(s string) string {
	newword := ""
	offsetn := int('z') - int('a') + 1
	offsetc := int('Z') - int('A') + 1
	for i := 0; i < len(s); i++ {
		n := s[i]
		if n >= 'a' && n <= 'z' {
			if n+14 > 'z' {
				newword = newword + string(n+14-byte(offsetn))
			} else {
				newword = newword + string(n+14)
			}
		} else if n >= 'A' && n <= 'Z' {
			if n+14 > 'Z' {
				newword = newword + string(n+14-byte(offsetc))
			} else {
				newword = newword + string(n+14)
			}
		} else {
			newword = newword + string(n)
		}

	}
	return newword
}
