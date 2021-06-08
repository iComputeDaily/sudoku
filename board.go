package main

import (
	"fmt"
	"math/rand"
)

type board interface {
	NumCells() int
	Cell(num int) *int
	NumGroups() int
	GroupSize() int
	Group(num int) []*int
	Clone() board
}

// Generate generates a valid sudoku board with blank spaces and only one solution
func Generate(b board) {
	Fill(b, func() bool {
		return false
	})

	// Remove numbers while keeping unique solution
	for nonBlanks := numNonBlank(b); nonBlanks >= 30; {
		// Get random cell
		cell := randNonBlank(b)

		// remove num while keeping backup
		backup := *cell
		*cell = 0

		if !hasOneSolution(b) {
			fmt.Println("reverting cell")
			*cell = backup
			continue
		}
		fmt.Println(nonBlanks)
		nonBlanks--
	}

	return
}

// hasOneSolution returns true if the puzzle has one solution
func hasOneSolution(b board) bool {
	// Create a deep copy of the board to make shure that we don't modify the callers board
	b = b.Clone()

	var numSolutions int

	Fill(b, func() bool {
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

// randNonBlank returns a pointer to a random non blank cell
func randNonBlank(b board) *int {
	cellNum := rand.Intn(numNonBlank(b) - 1)

	nonBlanks := 0 // Counts the number of non blank cells encountered

	// For all cells
	for i := 0; i < b.NumCells(); i++ {
		// Get the cell
		cell := b.Cell(i)

		// Do nothing if blank
		if *cell == 0 {
			continue
		}
		// Set return values if it is chosen cell
		if nonBlanks == cellNum {
			return cell
		}
		// Seen one more non blank cell
		nonBlanks++
	}
	return nil
}

// numNonBlank counts the number of non blank cells in a board
func numNonBlank(b board) (counter int) {
	for i := 0; i < b.NumCells(); i++ {
		if *b.Cell(i) != 0 {
			counter++
		}
	}
	return
}

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
