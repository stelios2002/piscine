package piscine

func Join(strs []string, sep string) string {
	newstring := ""
	for i := 0; i < len(strs); i++ {
		if i+1 < len(strs) {
			newstring = newstring + strs[i] + sep
		} else {
			newstring = newstring + strs[i]
		}
	}
	return newstring
}
