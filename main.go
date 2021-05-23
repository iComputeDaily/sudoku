package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SquareBoard are the actual sudoku numbers for a square board
type SquareBoard [9][9]int

// String prints the board
func (b *SquareBoard) String() (boardString string) {
	boardString += fmt.Sprintln(starterLine)

	for y := 0; y < 17; y++ { // for all rows
		switch {
		case y%2 == 0: // Number rows
			boardString += fmt.Sprint("│") // Leftmost wall

			for x := 0; x < 9; x++ {
				if x == 2 || x == 5 { // Number segment on group border
					boardString += fmt.Sprintf(" %-2d┃", b[x][y/2])
				} else { // Normal number segment
					boardString += fmt.Sprintf(" %-2d│", b[x][y/2])
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

// Group is a group of cells(exaple: row, collom, square)
type Group []*int // Pointers to board, not actual numbers

// Group returns the group of the number groupNum
func (b *SquareBoard) Group(groupNum int) (group Group) {
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

func main() {
	rand.Seed(time.Now().UnixNano())

	board := new(SquareBoard)

	board.Fill(func(b SquareBoard) bool {
		board = &b
		return false
	})

	fmt.Println(board)

	for i := 0; i < 27; i++ {
		coolGroup := board.Group(i)

		fmt.Print("Group ", i, ":\n")
		fmt.Print("\tActual value: ", coolGroup, "\n")
		fmt.Print("\tCells: [")
		for i := 0; i < 9; i++ {
			fmt.Printf("%d, ", *coolGroup[i])
		}
		fmt.Print("]\n")
	}
}
