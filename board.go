package main

import (
//	"fmt"
//	"math/rand"
)

type board interface {
	NumCells() int
	Cell(num int) *int
	NumGroups() int
	GroupSize() int
	Group(num int) []*int
}

/*
// Generate generates a valid sudoku board with blank spaces and only one solution
func Generate(b board) {
	Fill(b, func() bool {
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
*/

// check checks the board returning true if it is a valid part of a sudoku solution
func checkBoard(b board) bool {
	for i := 0; i < b.NumGroups(); i++ {
		g := b.Group(i)

		if !checkGroup(g) {
			return false
		}
	}

	return true
}
