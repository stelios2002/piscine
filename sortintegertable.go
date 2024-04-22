package piscine

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
