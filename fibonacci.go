package piscine

var (
	total = 1
	prev  = 0
	temp  = 0
)

func Fibonacci(index int) int {
	total = 1
	prev = 0
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	}
	GoDeeper(index)

	return total
}

func GoDeeper(left int) {
	if left > 1 {
		left -= 1
		temp = total
		total = total + prev
		prev = temp
		GoDeeper(left)
	}
	return
}
