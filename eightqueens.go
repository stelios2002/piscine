package piscine

import (
	"github.com/01-edu/z01"
)

const (
	EMPTY = 0
	PAINT = 1
	QUEEN = 2
)

var (
	board        [8][8]int // main board
	queenCounter = 0
)

// Grid[2][2] y,x
//  __________
// `          `
// [0, 0, 0, 0]
// ,
// [0, 0, 0, 0]
// ,
// [0, 0, 0, 0]
// ,
// [0, 0, 0, 0]
// ,___________,
//
//
//
//
//

func EightQueens() {
	// fmt.Println("EightQueens")
	placeQueen(0, 0)
}

func placeQueen(x, y int) {
	// fmt.Println("placeQueen")
	for xi := x; xi < 8; xi++ {
		for yi := 0; yi < 8; yi++ {
			if board[yi][xi] == EMPTY {
				addQueen(xi, yi)
				if queenCounter == 8 {
					// solutionsFound += 1
					// fmt.Println("found!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! ", solutionsFound)
					for xx := 0; xx < 8; xx++ {
						for yy := 0; yy < 8; yy++ {
							if board[yy][xx] == QUEEN {
								z01.PrintRune(rune(yy + 1 + 48))
							}
						}
					}
					z01.PrintRune(rune('\n'))
				} else {
					placeQueen(xi, yi)
					// fmt.Println("path finished")
					// printBoard()
				}
				removeQueen(xi, yi)
			}
		}
	}
}

func addQueen(x, y int) {
	// fmt.Println("addQueen")
	queenCounter += 1
	board[y][x] = QUEEN
	updatePaint()
}

func removeQueen(x, y int) {
	// fmt.Println("removeQueen")
	board[y][x] = EMPTY
	queenCounter -= 1
	updatePaint()
}

func updatePaint() {
	// fmt.Println("updatePaint")
	for yi := 0; yi < 8; yi++ {
		for xi := 0; xi < 8; xi++ {
			if board[yi][xi] != QUEEN {
				board[yi][xi] = EMPTY
			}
		}
	}

	for yi := 0; yi < 8; yi++ {
		for xi := 0; xi < 8; xi++ {
			if board[yi][xi] == QUEEN {
				for tx := 0; tx < 8; tx++ { // horizontal
					if board[yi][tx] == EMPTY {
						board[yi][tx] = PAINT
					}
				}
				for ty := 0; ty < 8; ty++ { // vertical
					if board[ty][xi] == EMPTY {
						board[ty][xi] = PAINT
					}
				}

				for i := 1; i < 9; i++ { // top right
					newx := xi + i
					newy := yi + i
					if isValidPos(newx, newy) == true {
						if board[newy][newx] == EMPTY {
							board[newy][newx] = PAINT
						}
					}
				}
				for i := 1; i < 9; i++ { // top left
					newx := xi - i
					newy := yi + i
					if isValidPos(newx, newy) == true {
						if board[newy][newx] == EMPTY {
							board[newy][newx] = PAINT
						}
					}
				}
				for i := 1; i < 9; i++ { // bot right
					newx := xi + i
					newy := yi - i
					if isValidPos(newx, newy) == true {
						if board[newy][newx] == EMPTY {
							board[newy][newx] = PAINT
						}
					}
				}
				for i := 1; i < 9; i++ { // bot left
					newx := xi - i
					newy := yi - i
					if isValidPos(newx, newy) == true {
						if board[newy][newx] == EMPTY {
							board[newy][newx] = PAINT
						}
					}
				}
			}
		}
	}
}

func generatePaintBoard() [8][8]int { // returns the paintboard used to memoize
	// fmt.Println("generatePaintBoard")
	PaintBoard := [8][8]int{}
	for yi := 0; yi < 8; yi++ {
		for xi := 0; xi < 8; xi++ {
			if board[yi][xi] == QUEEN || board[yi][xi] == PAINT {
				PaintBoard[yi][xi] = PAINT
			}
		}
	}
	return PaintBoard
}

// func isNewState(paintboard [8][8]int) bool {
// 	// fmt.Println("isNewState")

// 	if _, ok := memoizer[paintboard]; ok {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func addNewState(paintboard [8][8]int) {
// 	memoizer[paintboard] = 0
// }

func isValidPos(x, y int) bool {
	// fmt.Println("isValidPos")
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

// memoize the paint locations (treat queen as paint)

// func printBoard() {
// 	for yi := 0; yi < 8; yi++ {
// 		fmt.Println(board[yi])
// 	}
// 	fmt.Println("!!")
// }
