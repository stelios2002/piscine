package main

import (
	"os"

	"github.com/01-edu/z01"
)

// clean args off first
// detect --insert=/-i= or --order/-o or --help/-h or no arg
func main() {
	order := false
	insertWhat := "" // what to insert to whatever we're inserting to, if we're inserting something to something else
	bodyStr := ""    // this is what the inserted thing will get inserted to
	if len(os.Args) > 1 {
		// if true {
		// args := []string{"--insert=q9vbh9m414f 4", "--order", "EL5F  yZTaznX"}
		args := os.Args[1:]
		for arg_i := 0; arg_i < len(args); arg_i++ {
			curr_arg := args[arg_i]
			if curr_arg == "-o" || curr_arg == "--order" {
				order = true
			} else if curr_arg == "-h" || curr_arg == "--help" {
				displayHelp()
			} else if len(curr_arg) >= 3 { // what happens if it's just "-i=" and nothing else
				if Index(curr_arg, "-i=") == 0 {
					// insert
					insertWhat = curr_arg[3:]
				} else if len(curr_arg) > 8 {
					if Index(curr_arg, "--insert=") == 0 {
						// insert
						insertWhat = curr_arg[9:]
					} else {
						bodyStr = curr_arg
					}
				} else {
					bodyStr = curr_arg
				}
			} else {
				bodyStr = curr_arg
			}
		}
	} else {
		displayHelp()
		z01.PrintRune(rune('\n'))
		return
	}

	bodyStr = bodyStr + insertWhat

	runeSlice := []rune(bodyStr)
	intSlice := []int{}
	for i := 0; i < len(runeSlice); i++ {
		intSlice = append(intSlice, int(runeSlice[i]))
	}

	if order == true {
		SortIntegerTable(intSlice)
	}
	for i := 0; i < len(runeSlice); i++ {
		z01.PrintRune(rune(intSlice[i]))
	}
	z01.PrintRune(rune('\n'))
}

func displayHelp() {
	elp := []rune("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.")
	for i := 0; i < len(elp); i++ {
		z01.PrintRune(elp[i])
	}
}

func Index(s string, toFind string) int {
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

// sloppy merge sort
var global_table []int // holds a copy of the argument given to use its scope for ease of access in other functions

func SortIntegerTable(table []int) {
	global_table = table
	divide(0, len(global_table)-1)
	// fmt.Println(global_table)
}

// takes the two indeces of a slice, if the region is not 2 or 3 it calls itself
// to "cut" the regoins again until the sizes are 2 or 3 and then calls the combine function,
// if it's 2 or 3 it does a quick manual sort and returns
func divide(L int, R int) {
	if R-L+1 != 2 && R-L+1 != 3 {
		// generating new regions to feed into itself until they're small enough
		NewR := (L + R) / 2
		divide(L, NewR)
		NewL := NewR + 1
		divide(NewL, R)

		combine(L, NewR, NewL, R)

	} else {
		// manually sorting the regions of size 2 or 3
		if R-L+1 == 2 {
			if global_table[L] > global_table[R] {
				SwapThings(&global_table[L], &global_table[R])
			}
		} else {
			if global_table[L] > global_table[L+1] {
				SwapThings(&global_table[L], &global_table[L+1])
			}
			if global_table[L+1] > global_table[R] {
				SwapThings(&global_table[L+1], &global_table[R])
			}
			if global_table[L] > global_table[L+1] {
				SwapThings(&global_table[L], &global_table[L+1])
			}

		}
	}
}

// vvLeft index of A, Right index of A and so on..
func combine(LA int, RA int, LB int, RB int) { // takes two pairs of int indeces that represent two adjacent and sorted arrays, and then it merges those two sorted arrays together to fill their original location

	temp_slice := []int{} // sloppy... generating a temporary slice to hold the left side of the two adjacent arrays so values don't get overwritten, the right side doesn't need this because it's possible for it's values to be ovewritten??? or is it???? hmmm...
	DLA := 0              // These indedes are used instead of LA RA
	DRA := 0
	for i := LA; i <= RA; i++ {
		temp_slice = append(temp_slice, global_table[i])
		DRA += 1
	}

	OgIndex := LA // OgIndex will be used to write directly on the original slice given

	size := (RA - LA + 1) + (RB - LB + 1)
	for i := 0; i < size; i++ {

		if DLA < DRA && LB <= RB {
			if temp_slice[DLA] < global_table[LB] {
				global_table[OgIndex] = temp_slice[DLA]

				DLA += 1
			} else {
				global_table[OgIndex] = global_table[LB]

				LB += 1
			}
		} else {
			if DLA < DRA {
				// TODO possible break here?
				global_table[OgIndex] = temp_slice[DLA]
				DLA += 1

			} else if LB <= RB {
				global_table[OgIndex] = global_table[LB]

				LB += 1
			} else {
				break
			}
		}
		OgIndex += 1

	}
}

func SwapThings(a *int, b *int) { // swaps the values of the referenced variables
	temp := *b
	*b = *a
	*a = temp
}
