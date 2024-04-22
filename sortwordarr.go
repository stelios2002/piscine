package piscine

func SortWordArr(a []string) {
	SortTable(a)
}

// sloppy merge sort
var global_table2 []string // holds a copy of the argument given to use its scope for ease of access in other functions

func SortTable(table []string) {
	global_table2 = table
	divide2(0, len(global_table2)-1)
	// fmt.Println(global_table2)
}

// takes the two indeces of a slice, if the region is not 2 or 3 it calls itself
// to "cut" the regoins again until the sizes are 2 or 3 and then calls the combine function,
// if it's 2 or 3 it does a quick manual sort and returns
func divide2(L int, R int) {
	if R-L+1 != 2 && R-L+1 != 3 {
		// generating new regions to feed into itself until they're small enough
		NewR := (L + R) / 2
		divide2(L, NewR)
		NewL := NewR + 1
		divide2(NewL, R)

		combine2(L, NewR, NewL, R)

	} else {
		// manually sorting the regions of size 2 or 3
		if R-L+1 == 2 {
			if global_table2[L] > global_table2[R] {
				SwapThings2(&global_table2[L], &global_table2[R])
			}
		} else {
			if global_table2[L] > global_table2[L+1] {
				SwapThings2(&global_table2[L], &global_table2[L+1])
			}
			if global_table2[L+1] > global_table2[R] {
				SwapThings2(&global_table2[L+1], &global_table2[R])
			}
			if global_table2[L] > global_table2[L+1] {
				SwapThings2(&global_table2[L], &global_table2[L+1])
			}

		}
	}
}

// vvLeft index of A, Right index of A and so on..
func combine2(LA int, RA int, LB int, RB int) { // takes two pairs of int indeces that represent two adjacent and sorted arrays, and then it merges those two sorted arrays together to fill their original location

	temp_slice := []string{} // sloppy... generating a temporary slice to hold the left side of the two adjacent arrays so values don't get overwritten, the right side doesn't need this because it's possible for it's values to be ovewritten??? or is it???? hmmm...
	DLA := 0                 // These indedes are used instead of LA RA
	DRA := 0
	for i := LA; i <= RA; i++ {
		temp_slice = append(temp_slice, global_table2[i])
		DRA += 1
	}

	OgIndex := LA // OgIndex will be used to write directly on the original slice given

	size := (RA - LA + 1) + (RB - LB + 1)
	for i := 0; i < size; i++ {

		if DLA < DRA && LB <= RB {
			if temp_slice[DLA] < global_table2[LB] {
				global_table2[OgIndex] = temp_slice[DLA]

				DLA += 1
			} else {
				global_table2[OgIndex] = global_table2[LB]

				LB += 1
			}
		} else {
			if DLA < DRA {
				// TODO possible break here?
				global_table2[OgIndex] = temp_slice[DLA]
				DLA += 1

			} else if LB <= RB {
				global_table2[OgIndex] = global_table2[LB]

				LB += 1
			} else {
				break
			}
		}
		OgIndex += 1

	}
}

func SwapThings2(a *string, b *string) { // swaps the values of the referenced variables
	temp := *b
	*b = *a
	*a = temp
}
