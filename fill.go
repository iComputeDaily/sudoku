package main

import (
	"math/rand"
)

// Fill fills empty cells randomly, calling the callback on every solution found
func Fill(b board, callback func() bool) {
	trySolution(b, 0, callback)
}

// trySolution recursivly checks cells for a solution,
// bactracting until all are checked. On finding a solution
// it calls callback, and terminates if callback returns false.
func trySolution(b board, cellNum int, callback func() bool) bool {
	cell := b.Cell(cellNum)

	// Skip if not empty
	if *cell != 0 {
		if !recurse(b, cellNum, callback) {
			return false
		}
		return true
	}

	// Make a list of posibilities
	numList := make([]int, b.GroupSize())

	for i, _ := range numList {
		numList[i] = i + 1
	}

	rand.Shuffle(len(numList), func(x, y int) { numList[x], numList[y] = numList[y], numList[x] })

	// Try to recurse for all posibilities
	for _, num := range numList {
		*cell = num

		if !recurse(b, cellNum, callback) {
			return false
		}
	}

	// All nums tried
	*cell = 0
	return true
}

// recurse calls the callbackon solutions and recurses or returns acordingly
func recurse(b board, cellNum int, callback func() bool) bool {
	// Board is invalid
	if !checkBoard(b) {
		return true
	}

	// Board is valid
	if cellNum == b.NumCells()-1 { // Board is a solution
		if !callback() {
			return false
		}
	} else { // Board is valid but not a solution
		if !trySolution(b, cellNum+1, callback) {
			return false
		}
	}

	return true
}
