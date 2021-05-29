package main

// Group is a group of cells(exaple: row, collom, square)
type Group []*int // Pointers to board, not actual numbers

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
