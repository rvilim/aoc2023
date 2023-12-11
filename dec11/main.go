package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(filepath string) [][]rune {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result [][]rune

	for scanner.Scan() {
		lineText := scanner.Text()
		l := []rune(lineText)

		result = append(result, l)
	}

	return result
}

func rowEmpty(universe [][]rune, row int) bool {
	for _, col := range universe[row] {
		if col != '.' {
			return false
		}
	}
	return true
}

func colEmpty(universe [][]rune, col int) bool {
	for rowNum, _ := range universe {
		if universe[rowNum][col] != '.' {
			return false
		}
	}
	return true
}
func expand(universe [][]rune) [][]rune {
	var newUniverse [][]rune
	for rowNum, row := range universe {
		if rowEmpty(universe, rowNum) {
			newUniverse = append(newUniverse, row)
		}
		newUniverse = append(newUniverse, row)
	}

	// We are looping down from the top because appending muddles up our array indices for all indices > than our current one
	for colNum := len(universe[0]) - 1; colNum >= 0; colNum-- {
		if colEmpty(universe, colNum) {
			for rowNum, _ := range newUniverse {
				newRow := newUniverse[rowNum][:colNum]
				newRow = append(newRow, '.')
				newRow = append(newRow, newUniverse[rowNum][colNum:]...)
				newUniverse[rowNum] = newRow
			}
		}

	}
	return newUniverse
}

func printUniverse(universe [][]rune) {
	for _, row := range universe {
		fmt.Println(string(row))
	}
}

type Coord struct {
	row int
	col int
}

func countPath(universe [][]rune, expansionFactor int) int {
	galaxies := []Coord{}

	for rowNum, row := range universe {
		for colNum, char := range row {
			if char == '#' {
				galaxies = append(galaxies, Coord{row: rowNum, col: colNum})
			}
		}
	}

	total := 0
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			for rowNum := galaxy1.row; rowNum < galaxy2.row; rowNum++ {
				if rowEmpty(universe, rowNum) {
					total += expansionFactor
				} else {
					total += 1
				}
			}

			for colNum := galaxy1.col; colNum < galaxy2.col; colNum++ {
				if colEmpty(universe, colNum) {
					total += expansionFactor
				} else {
					total += 1
				}
			}
		}
	}
	return total
}
func q1(universe [][]rune) int {
	return countPath(universe, 2)
	// universe = expand(universe)
	// // printUniverse(universe)
	// galaxies := []Coord{}

	// for rowNum, row := range universe {
	// 	for colNum, char := range row {
	// 		if char == '#' {
	// 			galaxies = append(galaxies, Coord{row: rowNum, col: colNum})
	// 		}
	// 	}
	// }

	// total := 0
	// for _, galaxy1 := range galaxies {
	// 	for _, galaxy2 := range galaxies {
	// 		total += Abs(galaxy1.row-galaxy2.row) + Abs(galaxy1.col-galaxy2.col)
	// 	}
	// }
	// return total / 2
}

func q2(universe [][]rune) int {

	return countPath(universe, 1000000)
}
func main() {
	universe := parse("dec11/data/input.txt")
	println(q1(universe))

	universe = parse("dec11/data/input.txt")
	println(q2(universe))
}
