package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	board := Generate()

	fmt.Println(board)

	for i := 0; i < 27; i++ {
		coolGroup := board.Group(i)

		fmt.Print("Group ", i, ":\n")
		fmt.Print("\tActual value: ", coolGroup, "\n")
		fmt.Print("\tCells: [")
		for i := 0; i < 9; i++ {
			fmt.Printf("%d, ", *coolGroup[i])
		}
		fmt.Print("]\n")
	}
}
