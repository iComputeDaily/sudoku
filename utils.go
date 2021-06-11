package sudoku

import "math/rand"

// removeNums recursivly removes numbers from the board returning false if it did not reach targetClues
func removeNums(b board, targetClues int) bool {
	// Get list of all non blank cells
	cellList := make([]*int, 0, b.NumCells())

	// For all cells
	for i := 0; i < b.NumCells(); i++ {
		cell := b.Cell(i)

		// Add to the list if not blank
		if *cell != 0 {
			cellList = append(cellList, cell)
		}
	}
	// Randomize
	rand.Shuffle(len(cellList), func(x, y int) { cellList[x], cellList[y] = cellList[y], cellList[x] })

	// Try blanking all cells
	for _, cell := range cellList {
		// Keep backup and make 0
		backup := *cell
		*cell = 0

		// Revert if multiple solutions
		if !hasOneSolution(b) {
			*cell = backup
			continue
		}

		// Return sucsess if blanked enouph cells
		if numNonBlank(b) == targetClues {
			return true
		}

		// Otherwise continue trying
		if removeNums(b, targetClues) {
			return true
		}

		*cell = backup
	}

	return false
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

	// Return false if multiple solutions found
	if numSolutions >= 2 {
		return false
	}

	return true
}

// randNonBlank returns a pointer to a random non blank cell
func randNonBlank(b board) *int {
	// Pick a random non blank cell number
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
		// Return if it is the randomly chosen cell number
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

		// if any group fails whole board fails
		if !checkGroup(g) {
			return false
		}
	}

	return true
}

// check returns true if the group has no duplicate numbers
func checkGroup(g []*int) bool {
	numIndex := make(map[int]bool)

	// For all nums in group
	for i := 0; i < len(g); i++ {
		// If number previously seen it is a duplicate
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
