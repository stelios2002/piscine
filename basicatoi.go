package piscine

func BasicAtoi(s string) int {
	temp := 0
	sig := 1
	for i := len(s) - 1; i >= 0; i-- {
		temp += sig * (int(rune(s[i]) - 48))
		sig *= 10
	}
	return temp
}
