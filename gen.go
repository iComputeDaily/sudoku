package main

import (
	"fmt"
	"math/rand"
)

func (b *SquareBoard) Generate() {
	ok := b.tryNums(0, 0)
	if ok == false {
		fmt.Println("Somehow tryNums failed")
	}
}

// tryNums trys all the numbers for the cell passed in, and
// calls itself on the next cell recursivly. Returns true if it
// reaches the last cell, and false if it reaches a dead end.
func (b *SquareBoard) tryNums(x, y int) bool {
	// Make a randomized list on numbers
	numList := make([]int, 9)
	for i := 0; i < len(numList); i++ {
		numList[i] = i + 1
	}
	rand.Shuffle(len(numList), func(x, y int) { numList[x], numList[y] = numList[y], numList[x] })

	// Try all nums
	for _, num := range numList {
		b[x][y] = num

		if b.check() {
			switch {
			// Reached the end
			case x == 8 && y == 8:
				return true

			// Continue on current line
			case x < 8:
				if b.tryNums(x+1, y) {
					return true
				}

			// Move to next line
			case x == 8:
				if b.tryNums(0, y+1) {
					return true
				}
			}
		}
	}

	// All nums tried, no solution found
	b[x][y] = 0
	return false
}
