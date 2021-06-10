package sudoku

type board interface {
	NumCells() int
	Cell(num int) *int
	NumGroups() int
	GroupSize() int
	Group(num int) []*int
	Clone() board
}

// Generate generates a valid sudoku board with blank spaces and only one solution
func Generate(b board, numGivens int) {
	// Randomly populate board
	Fill(b, func() bool {
		return false
	})

	// Remove numbers while keeping unique solution
	for nonBlanks := numNonBlank(b); nonBlanks >= numGivens; {
		// Get random cell
		cell := randNonBlank(b)

		// remove num while keeping backup
		backup := *cell
		*cell = 0

		if !hasOneSolution(b) {
			*cell = backup
			continue
		}
		nonBlanks--
	}

	return
}

// Solve returns true if the board was solved, and false if the board was not
func Solve(b board) (solved bool) {
	Fill(b, func() bool {
		solved = true
		return false
	})
	return solved
}
