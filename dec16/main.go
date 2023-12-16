package main

import (
	"bufio"
	"os"
)

func parse(filepath string) [][]rune {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type coord struct {
	row int
	col int
}

type beamVec struct {
	coord
	direction Direction
}

func moveCoord(initCoord coord, direction Direction) coord {
	switch direction {
	case Left:
		return coord{row: initCoord.row, col: initCoord.col - 1}
	case Right:
		return coord{row: initCoord.row, col: initCoord.col + 1}
	case Up:
		return coord{row: initCoord.row - 1, col: initCoord.col}
	case Down:
		return coord{row: initCoord.row + 1, col: initCoord.col}
	default:
		panic("Invalid direction")
	}
}

func step(grid [][]rune, stepCache map[beamVec]bool, beam beamVec) {
	if beam.col == len(grid[0]) || beam.col < 0 || beam.row < 0 || beam.row == len(grid) {
		return
	}

	var newDirections []Direction
	switch grid[beam.row][beam.col] {
	case '.':
		{
			newDirections = []Direction{beam.direction}
		}
	case '/':
		{
			switch beam.direction {
			case Left:
				newDirections = []Direction{Down}
			case Right:
				newDirections = []Direction{Up}
			case Down:
				newDirections = []Direction{Left}
			case Up:
				newDirections = []Direction{Right}
			default:
				println("Error, unknown direction")
			}
		}
	case '\\':

		{
			switch beam.direction {
			case Left:
				newDirections = []Direction{Up}
			case Right:
				newDirections = []Direction{Down}
			case Down:
				newDirections = []Direction{Right}
			case Up:
				newDirections = []Direction{Left}
			default:
				println("Error, unknown direction")
			}
		}
	case '-':
		{
			switch beam.direction {
			case Left, Right:
				newDirections = []Direction{beam.direction}
			case Up, Down:
				newDirections = []Direction{Left, Right}
			default:
				println("Error, unknown direction")
			}
		}
	case '|':
		{
			switch beam.direction {
			case Left, Right:
				newDirections = []Direction{Up, Down}
			case Up, Down:
				newDirections = []Direction{beam.direction}
			default:
				println("Error, unknown direction")
			}
		}
	}

	for _, newDirection := range newDirections {
		newBeamVec := beamVec{
			coord:     moveCoord(beam.coord, newDirection),
			direction: newDirection,
		}
		_, exists := stepCache[newBeamVec]
		stepCache[newBeamVec] = true
		if !exists {
			step(grid, stepCache, newBeamVec)
		}
	}
}

func score(stepCache map[beamVec]bool, grid [][]rune) int {
	total := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			_, upOk := stepCache[beamVec{coord{row: row, col: col}, Up}]
			_, downOk := stepCache[beamVec{coord{row: row, col: col}, Down}]
			_, leftOk := stepCache[beamVec{coord{row: row, col: col}, Left}]
			_, rightOk := stepCache[beamVec{coord{row: row, col: col}, Right}]

			if upOk || downOk || leftOk || rightOk {
				total += 1
			}
		}
	}
	return total
}
func printEnergized(stepCache map[beamVec]bool, grid [][]rune) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			_, upOk := stepCache[beamVec{coord{row: row, col: col}, Up}]
			_, downOk := stepCache[beamVec{coord{row: row, col: col}, Down}]
			_, leftOk := stepCache[beamVec{coord{row: row, col: col}, Left}]
			_, rightOk := stepCache[beamVec{coord{row: row, col: col}, Right}]

			if upOk || downOk || leftOk || rightOk {
				print("#")
			} else {
				print(".")
			}
		}
		println("")
	}
}
func q1(grid [][]rune) int {
	stepCache := make(map[beamVec]bool)
	initVec := beamVec{coord{row: 0, col: 0}, Right}
	stepCache[initVec] = true
	step(grid, stepCache, initVec)

	return score(stepCache, grid)

}

func processEdge(grid [][]rune, startRow, startCol int, direction Direction, maxScore *int) {
	stepCache := make(map[beamVec]bool)

	initVec := beamVec{coord{row: startRow, col: startCol}, direction}
	stepCache[initVec] = true

	step(grid, stepCache, initVec)

	curScore := score(stepCache, grid)
	if curScore > *maxScore {
		*maxScore = curScore
	}
}

func q2(grid [][]rune) int {
	var maxScore int

	// top and bottom
	for col := 0; col < len(grid[0]); col++ {
		processEdge(grid, 0, col, Down, &maxScore)
		processEdge(grid, len(grid)-1, col, Up, &maxScore)

	}
	//  left and right
	for row := 0; row < len(grid); row++ {
		processEdge(grid, row, 0, Right, &maxScore)
		processEdge(grid, row, len(grid[0])-1, Left, &maxScore)
	}
	return maxScore
}

func main() {
	l := parse("dec16/data/input.txt")

	println(q1(l))
	println(q2(l))
}
