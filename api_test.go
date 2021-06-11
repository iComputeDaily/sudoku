package sudoku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	board := new(SquareBoard)

	Generate(board, 30)

	fmt.Println(board)
}

func TestSolveUnsolvable(t *testing.T) {
	board := &SquareBoard{
		[9]int{5, 9, 0, 3, 0, 0, 5, 0, 0},
		[9]int{0, 0, 0, 0, 8, 1, 0, 6, 5},
		[9]int{0, 1, 0, 0, 0, 0, 3, 0, 0},
		[9]int{1, 6, 0, 0, 3, 2, 9, 0, 0},
		[9]int{0, 4, 0, 0, 5, 8, 0, 0, 6},
		[9]int{0, 0, 0, 0, 9, 6, 0, 7, 0},
		[9]int{0, 0, 0, 0, 4, 0, 0, 0, 1},
		[9]int{2, 5, 0, 8, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 3, 2, 0, 4}}

	expected := &SquareBoard{
		[9]int{5, 9, 0, 3, 0, 0, 5, 0, 0},
		[9]int{0, 0, 0, 0, 8, 1, 0, 6, 5},
		[9]int{0, 1, 0, 0, 0, 0, 3, 0, 0},
		[9]int{1, 6, 0, 0, 3, 2, 9, 0, 0},
		[9]int{0, 4, 0, 0, 5, 8, 0, 0, 6},
		[9]int{0, 0, 0, 0, 9, 6, 0, 7, 0},
		[9]int{0, 0, 0, 0, 4, 0, 0, 0, 1},
		[9]int{2, 5, 0, 8, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 3, 2, 0, 4}}

	solved := Solve(board)

	if *board != *expected {
		t.Error("Solve function modified the board when there was no solution\n\nResult:\n", board, "\n\nExpected:\n", expected)
	}
	if solved == true {
		t.Error("Solve function returned true even though the board was unsolvable")
	}
}

func TestSolve(t *testing.T) {
	expected := &SquareBoard{
		[9]int{5, 9, 6, 3, 7, 4, 8, 1, 2},
		[9]int{4, 3, 2, 9, 8, 1, 7, 6, 5},
		[9]int{7, 1, 8, 6, 2, 5, 3, 4, 9},
		[9]int{1, 6, 7, 4, 3, 2, 9, 5, 8},
		[9]int{3, 4, 9, 7, 5, 8, 1, 2, 6},
		[9]int{8, 2, 5, 1, 9, 6, 4, 7, 3},
		[9]int{6, 8, 3, 2, 4, 7, 5, 9, 1},
		[9]int{2, 5, 4, 8, 1, 9, 6, 3, 7},
		[9]int{9, 7, 1, 5, 6, 3, 2, 8, 4}}

	var numFails int

	for i := 0; i < 3; i++ {
		board := &SquareBoard{
			[9]int{5, 9, 0, 3, 0, 0, 0, 0, 0},
			[9]int{0, 0, 0, 0, 8, 1, 0, 6, 5},
			[9]int{0, 1, 0, 0, 0, 0, 3, 0, 0},
			[9]int{1, 6, 0, 0, 3, 2, 9, 0, 0},
			[9]int{0, 4, 0, 0, 5, 8, 0, 0, 6},
			[9]int{0, 0, 0, 0, 9, 6, 0, 7, 0},
			[9]int{0, 0, 0, 0, 4, 0, 0, 0, 1},
			[9]int{2, 5, 0, 8, 0, 0, 0, 0, 0},
			[9]int{0, 0, 0, 0, 0, 3, 2, 0, 4}}
		solved := Solve(board)

		if *board != *expected || solved != true {
			numFails++
		}
		if *board != *expected {
			t.Log("Incorrect solution on try", i+1, "/", 3, "\n\nResult:\n", board)
		}
		if solved != true {
			t.Log("Solve returned false with solvable board on try", i+1, "/", 3)
		}
	}

	if numFails > 0 {
		t.Log("The test failed", numFails, "/", 3, "times")
		t.Log("Expected result was:\n", expected)
		t.Fail()
	}
}

func TestString(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	expected := `┌───┬───┬───┰───┬───┬───┰───┬───┬───┐
│   │ 3 │ 5 ┃ 4 │ 6 │   ┃   │ 7 │ 1 │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│   │   │   ┃   │   │ 1 ┃   │   │ 9 │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│   │   │ 37┃   │   │ 3 ┃   │ 8 │   │
┝━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┥
│ 2 │ 6 │   ┃   │   │ 4 ┃   │   │ 7 │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│   │   │ 1 ┃   │   │   ┃ 4 │   │ 2 │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│   │ 7 │ 8 ┃   │   │   ┃   │ 1 │   │
┝━━━┿━━━┿━━━╋━━━┿━━━┿━━━╋━━━┿━━━┿━━━┥
│   │   │   ┃ 5 │   │   ┃ 9 │ 6 │   │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│ 5 │   │ 3 ┃   │   │   ┃   │   │   │
├───┼───┼───╂───┼───┼───╂───┼───┼───┤
│   │ 8 │   ┃ 3 │ 7 │   ┃   │   │ 5 │
└───┴───┴───┸───┴───┴───┸───┴───┴───┘`

	board := &SquareBoard{
		[9]int{0, 0, 0, 2, 0, 0, 0, 5, 0},
		[9]int{3, 0, 0, 6, 0, 7, 0, 0, 8},
		[9]int{5, 0, 37, 0, 1, 8, 0, 3, 0},
		[9]int{4, 0, 0, 0, 0, 0, 5, 0, 3},
		[9]int{6, 0, 0, 0, 0, 0, 0, 0, 7},
		[9]int{0, 1, 3, 4, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 4, 0, 9, 0, 0},
		[9]int{7, 0, 8, 0, 0, 1, 6, 0, 0},
		[9]int{1, 9, 0, 7, 2, 0, 0, 0, 5}}

	out := fmt.Sprint(board)

	if expected != out {
		t.Log("Expected:\n", expected)
		t.Log("Result:\n", out)
		t.Fail()
	}
}
