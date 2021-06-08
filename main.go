package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := new(SquareBoard)

	Generate(board)

	fmt.Println(board)
}
