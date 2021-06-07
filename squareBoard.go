package main

import (
	"fmt"
)

// SquareBoard are the actual sudoku numbers for a square board
type SquareBoard [9][9]int

// String prints the board
func (b SquareBoard) String() (boardString string) {
	boardString += fmt.Sprintln(starterLine)

	for y := 0; y < 17; y++ { // for all rows
		switch {
		case y%2 == 0: // Number rows
			boardString += fmt.Sprint("│") // Leftmost wall

			for x := 0; x < 9; x++ {
				// Make blank cell for 0
				if x == 2 || x == 5 { // Number segment on group border
					if b[x][y/2] == 0 {
						boardString += fmt.Sprintf("   ┃")
					} else {
						boardString += fmt.Sprintf(" %-2d┃", b[x][y/2])
					}
				} else { // Normal number segment
					if b[x][y/2] == 0 {
						boardString += fmt.Sprintf("   │")
					} else {
						boardString += fmt.Sprintf(" %-2d│", b[x][y/2])
					}
				}
			}

		default: // line rows
			if y%6 == 5 {
				boardString += fmt.Sprint(boldLine)
			} else {
				boardString += fmt.Sprint(normalLine)
			}
		}
		boardString += fmt.Sprintf("\n")
	}
	boardString += fmt.Sprintln(endingLine)

	return
}

// NumCells returns the number of cells
func (b SquareBoard) NumCells() int {
	return 81
}

// Cell returns a pointer to the cell with number num
func (b *SquareBoard) Cell(num int) *int {
	return &b[num%9][num/9]
}

// NumGroups returns the number of groups(like rows, couloms, and squares) in the board
func (b SquareBoard) NumGroups() int {
	return 27
}

// GroupSize returns the number of cells per group
func (b SquareBoard) GroupSize() int {
	return 9
}

// Group returns the group of the number groupNum
func (b *SquareBoard) Group(groupNum int) (group []*int) {
	group = make([]*int, 9)

	switch {
	// rows
	case groupNum >= 0 && groupNum <= 8:
		for i := range group {
			group[i] = &b[i][groupNum]
		}
	// colomns
	case groupNum >= 9 && groupNum <= 17:
		for i := range group {
			group[i] = &b[groupNum-9][i]
		}
	// squares
	case groupNum >= 18 && groupNum <= 26:
		var i = 0
		for y := 0; y <= 2; y++ {
			for x := 0; x <= 2; x++ {
				group[i] = &b[((groupNum-18)%3*3)+x][(groupNum-18)/3*3+y]
				i++
			}
		}
	// invalid
	default:
		return
	}
	return
}
