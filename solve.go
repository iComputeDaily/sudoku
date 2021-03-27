package main

// checkDuplicates returns false if there are duplicates, and true if there aren't.
func (g Group) checkDuplicates() bool {
	numTable := make(map[int]bool)

	for i := 0; i < len(g); i++ {
		_, exists := numTable[*g[i]]
		if exists {
			return false
		}
		if *g[i] != 0 {
			numTable[*g[i]] = true
		}
	}

	return true
}

// Checks the whole board. Returns true if it is solved, and false if it isn't.
func (b *SquareBoard) check() bool {
	for i := 0; i < 27; i++ {
		g := b.Group(i)

		if !g.checkDuplicates() {
			return false
		}
	}

	return true
}
