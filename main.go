package main

import (
	fmt "fmt"
	"math/rand"
	"time"
)

var n int
var seed int64
var m int

type space [][]string

func getCoordinates(position int, offset int, size int) int {
	if (position + offset) < 0 {
		return size - 1
	} else if (position + offset) > (size - 1) {
		return 0
	} else {
		return position + offset
	}
}
func cellsFuture(universe space, position [2]int, size int) bool {
	var isNeighbourAlive int
	positions := []int{-1, 0, 1}
	for _, row := range positions {
		var rowCoordinates int = getCoordinates(position[0], row, size)
		for _, column := range positions {
			var columnCoordinates int = getCoordinates(position[1], column, size)
			if row != 0 || column != 0 {
				if universe[rowCoordinates][columnCoordinates] == "O" {
					isNeighbourAlive += 1
				}
			}
		}
	}
	if isNeighbourAlive >= 2 && isNeighbourAlive <= 3 && universe[position[0]][position[1]] == "O" {
		return true
	} else if isNeighbourAlive == 3 && universe[position[0]][position[1]] == " " {
		return true
	} else {
		return false
	}
}

func printUniverse(universe space, alive int, generation int) {
	fmt.Printf("Generation #%d\n", generation)
	fmt.Printf("Alive: %d\n", alive)

	for row := range universe {
		for column := range (universe)[row] {
			fmt.Print(universe[row][column])
		}
		fmt.Print("\n")
	}
}

func main() {
	fmt.Scan(&n)

	//rand.Seed(seed)

	generation := 1
	alive := 0

	var universe space = make(space, n)

	// define universe
	for i := 0; i < n; i++ {
		universe[i] = make([]string, n)
		for j := 0; j < n; j++ {
			if rand.Intn(2) == 1 {
				universe[i][j] = "O"
				alive++
			} else {
				universe[i][j] = " "
			}

		}
		// fmt.Println(universe[i])
	}
	printUniverse(universe, alive, generation)
	alive = 0
	generation++
	for {
		nextGenUniverse := make(space, n)
		time.Sleep(500 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		// fmt.Printf("Generation: %d\n", generation)
		for i := 0; i < n; i++ {
			nextGenUniverse[i] = make([]string, n)
			for j := 0; j < n; j++ {
				if cellsFuture(universe, [2]int{i, j}, n) == true {
					nextGenUniverse[i][j] = "O"
					alive++
				} else {
					nextGenUniverse[i][j] = " "
				}
			}
		}
		copy(universe, nextGenUniverse)
		nextGenUniverse = nil
		printUniverse(universe, alive, generation)
		alive = 0
		generation++
		if generation == 40 {
			break
		}
	}
	// printUniverse(universe)
}
