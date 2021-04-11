package main

import (
// "math/rand"
)

// checkDuplicates returns false if there are duplicates, and true if there aren't.
func (g Group) check() bool {
	numIndex := make(map[int]bool)

	for i := 0; i < len(g); i++ {
		// If number seen it isn't a solution
		_, exists := numIndex[*g[i]]
		if exists {
			return false
		}

		// If zero sulution not found
		if *g[i] == 0 {
			return false
		}

		// Mark number as seen
		numIndex[*g[i]] = true
	}

	return true
}

// Checks the whole board. Returns true if it is solved, and false if it isn't.
func (b *SquareBoard) check() bool {
	for i := 0; i < 27; i++ {
		g := b.Group(i)

		if !g.check() {
			return false
		}
	}

	return true
}

// trySolution recursivly checks cells for a solution, bactracting until all are checked
func (b *SquareBoard) trySolution(x, y, numSolutions int) bool {
	if b[x][y] != 0 {
		switch {
		// Reached the end
		case b.check():
			if numSolutions+1 >= 2 {
				return true
			}

		// Continue on current line
		case x < 8:
			return b.trySolution(x+1, y, numSolutions)

		// Move to next line
		case x == 8:
			return b.trySolution(0, y+1, numSolutions)
		}
	}

	// Make a list of posibilities
	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, num := range numList {
		b[x][y] = num

	}
}
