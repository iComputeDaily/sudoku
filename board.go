package main

import (
	"fmt"
	"math/rand"
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

// Generate generates a valid sudoku board with blank spaces and only one solution
func Generate() (b SquareBoard) {
	b.Fill(func(board SquareBoard) bool {
		b = board
		return false
	})

	// Remove numbers while keeping unique solution
	for nonBlanks := b.numNonBlank(); !(nonBlanks <= 30); {
		// Get random cell
		_, _, cell := b.randNonBlank()

		// remove num while keeping backup
		backup := *cell
		*cell = 0

		if !b.hasOneSolution() {
			*cell = backup
			continue
		}
		fmt.Println(nonBlanks)
		nonBlanks--
	}

	return
}

// hasOneSolution returns true if the puzzle has one solution
func (b SquareBoard) hasOneSolution() bool {
	var numSolutions int

	b.Fill(func(board SquareBoard) bool {
		// Stop searching if solution already found
		if numSolutions == 1 {
			numSolutions++
			return false
		}

		// Otherwise keep searching
		numSolutions++
		return true
	})

	if numSolutions >= 2 {
		return false
	}

	return true
}

// randNonBlank returns a random non blank cell
func (b *SquareBoard) randNonBlank() (x int, y int, cell *int) {
	cellNum := rand.Intn(b.numNonBlank() - 1)

	nonBlanks := 0 // Counts the number of non blank cells encountered

	// For all cells
	for h := 0; h < 9; h++ {
		for v := 0; v < 9; v++ {
			// Do nothing if blank
			if b[h][v] == 0 {
				continue
			}
			// Set return values if it is chosen cell
			if nonBlanks == cellNum {
				x = h
				y = v
				cell = &b[h][v]
				return
			}
			// Seen one more non blank cell
			nonBlanks++
		}
	}
	return
}

// numNonBlank counts the number of non blank cells in a board
func (b SquareBoard) numNonBlank() (counter int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if b[x][y] != 0 {
				counter++
			}
		}
	}
	return
}

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

// check checks the board returning true if it is a valid part of a sudoku solution
func (b *SquareBoard) check() bool {
	for i := 0; i < 27; i++ {
		g := b.Group(i)

		if !g.check() {
			return false
		}
	}

	return true
}
