package main

import (
	"math/rand"
)

// Generate generates a valid sudoku board with blank spaces and only one solution
func Generate() (b SquareBoard) {
	b.Fill(func(board SquareBoard) bool {
		b := board
		return false
	})
}

// Fill fills empty cells randomly, calling the callback on every solution found
func (b SquareBoard) Fill(callback func(board SquareBoard) bool) {
	b.trySolution(0, 0, callback)
}

// check returns true if the group is part of a valid sudoku solution
func (g Group) check() bool {
	numIndex := make(map[int]bool)

	// For all nums in group
	for i := 0; i < len(g); i++ {
		// If number seen it isn't a solution
		_, exists := numIndex[*g[i]]
		if exists {
			return false
		}

		// If cell is not empty(0), mark as seen
		if *g[i] != 0 {
			numIndex[*g[i]] = true
		}
	}

	return true
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

// recurse calls the callbackon solutions and recurses or returns acordingly
func (b SquareBoard) recurse(x, y int, callback func(board SquareBoard) bool) bool {
	// Board is invalid
	if !b.check() {
		return true
	}

	// Board is valid
	if x == 8 && y == 8 { // Board is a solution
		if !callback(b) {
			return false
		}
	} else { // Board is valid but not a solution
		switch {
		case x == 8:
			if !b.trySolution(0, y+1, callback) {
				return false
			}

		default:
			if !b.trySolution(x+1, y, callback) {
				return false
			}
		}
	}

	return true
}

// trySolution recursivly checks cells for a solution,
// bactracting until all are checked. On finding a solution
// it calls callback, and terminates if callback returns false.
func (b SquareBoard) trySolution(x, y int, callback func(board SquareBoard) bool) bool {
	// Skip if already filled in
	if b[x][y] != 0 {
		if !b.recurse(x, y, callback) {
			return false
		}
		return true
	}

	// Make a list of posibilities
	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numList), func(x, y int) { numList[x], numList[y] = numList[y], numList[x] })

	// Try to recurse for all posibilities
	for _, num := range numList {
		b[x][y] = num

		if !b.recurse(x, y, callback) {
			return false
		}
	}

	// All nums tried
	b[x][y] = 0
	return true
}
