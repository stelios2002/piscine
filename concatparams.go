package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			result = result + args[i]
		} else {
			result = result + args[i] + "\n"
		}
	}
	return result
}
