package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := new(SquareBoard)

	Fill(board, func() bool {
		return false
	})

	fmt.Println(board)
}
