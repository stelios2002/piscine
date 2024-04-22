package main

import "os"

func main() {
	vala := 0
	valb := 0
	res := 0

	symb := ""
	args := os.Args[1:]

	if len(args) == 3 { //|| true {
		if IsNumeric(args[0]) == false || IsNumeric(args[2]) == false {
			return
		}
		isok := true
		vala, isok = Atoi3(args[0])
		valb, isok = Atoi3(args[2])
		if isok == false {
			return
		}
		// vala = -927
		// valb = 868
		//"-927" "/" "868"
		symb = args[1]
		// symb = "/"
		// do stuff
		//+, -, /, *, %
		switch symb {
		case "+":

			res = vala + valb
			if res < vala && res < valb {
				return
			}
			if res-vala != valb {
				// overflow!
				return
			} else {
				// ok!
			}

		case "-":
			res = vala - valb
			if res > vala && res > valb {
				return
			}
			if res+valb != vala {
				// overflow!
				return
			} else {
				// ok!
			}
		case "/":
			if valb == 0 {
				// cant div by 0
				os.Stdout.WriteString("No division by 0\n")
				return
			} else {
				res = vala / valb
				// no need to check for overflow
			}

		case "*":
			res = vala * valb
			if res/valb != vala {
				// overflow!
				return
			} else {
				// ok!
			}

		case "%":
			if valb == 0 {
				// cant mod by 0
				os.Stdout.WriteString("No modulo by 0\n")
				return
			} else {
				res = vala % valb
			}

		default:
			return
		}
		result = ""
		StringNbr(res)
		os.Stdout.WriteString(result)
		os.Stdout.WriteString("\n")

	} else if len(args) > 3 {
		// too many args
	} else {
		// too few args
	}
}

func Atoi3(s string) (int, bool) {
	var valids2 string
	if len(s) == 0 {
		return 0, false
	}
	symbol_count := 0
	valids2 = "0123456789-+"
	for si := 0; si < len(s); si++ {
		FoundMatch := false
		for vi := 0; vi < len(valids2); vi++ {
			if valids2[vi] == s[si] {
				if s[si] == valids2[10] || s[si] == valids2[11] {
					symbol_count++
					if symbol_count > 1 || si > 0 {
						return 0, false
					}
				}

				FoundMatch = true
				break
			}
		}
		if FoundMatch == false {
			return 0, false
		}
	}
	is_neg := false
	if s[0] == valids2[10] {
		is_neg = true
	}
	oldtemp := 0
	temp := 0
	sig := 1

	for i := len(s) - 1; i >= symbol_count; i-- {
		oldtemp = temp
		temp += sig * (int(rune(s[i]) - 48))
		if temp-sig*(int(rune(s[i])-48)) != oldtemp {
			// overflow!
			return 0, false
		}
		sig *= 10
	}

	if is_neg == true {
		temp = -temp
	}

	return temp, true
}

var result string = ""

func StringNbr(n int) {
	temp := n

	if n == -9223372036854775808 {
		temp = 8
		n += 1
		n = -n
		result = result + "-"

	} else {

		if n < 0 {
			result = result + "-"
			n = -n
		}
		temp = n % 10
	}

	n /= 10

	if n > 0 {
		StringNbr(n)
	}

	result = result + string(rune(temp+48))

	return
}

func IsNumeric(s string) bool {
	runeslice := []rune(s)
	for i := 0; i < len(runeslice); i++ {
		if int(runeslice[i]) >= 48 && int(runeslice[i]) <= 57 || int(runeslice[i]) == '-' || int(runeslice[i]) == '+' {
		} else {
			return false
		}
	}
	return true
}
