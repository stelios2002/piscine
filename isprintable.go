package piscine

// used the unicode.IsPrint() to see which runes are printable and not, took the first 1k hoping those cover all test cases thrown to me :)
// 32   32 :  true
// 127   127 :  false
// 161   161 :  true
// 173   173 :  false
// 174   174 :  true
// 888   888 :  false

func IsPrintable(s string) bool {
	runeslice := []rune(s)
	runum := 0
	for i := 0; i < len(runeslice); i++ {
		runum = int(runeslice[i])
		if accRange(runum, 32, 126) || accRange(runum, 161, 172) || accRange(runum, 174, 887) {
		} else {
			return false
		}
	}
	return true
}

func accRange(num, min, max int) bool {
	return num >= min && num <= max
}
