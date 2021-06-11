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
	for {
		board := b.Clone()

		// Randomly populate board
		Fill(board, func() bool {
			return false
		})

		// Randomly remove numbers
		if removeNums(board, numGivens) {
			b = board
			break
		}
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
