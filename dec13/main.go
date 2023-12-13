package main

import (
	"bufio"
	"os"
	"slices"
)

func parse(filepath string) [][][]rune {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var tempResult [][]rune
	var result [][][]rune
	for scanner.Scan() {
		lineText := scanner.Text()

		if lineText == "" {
			result = append(result, tempResult)
			tempResult = [][]rune{}
		} else {
			tempResult = append(tempResult, []rune(lineText))
		}

	}
	result = append(result, tempResult)

	return result
}

func extractCol(input [][]rune, col int) []rune {
	var s []rune

	for _, row := range input {
		s = append(s, rune(row[col]))
	}
	return s
}
func transpose(input [][]rune) [][]rune {
	var s [][]rune
	for col := 0; col < len(input[0]); col++ {
		s = append(s, extractCol(input, col))
	}
	return s
}
func testReflectCol(input [][]rune, col int) bool {
	col += 1
	for i := 1; col-i >= 0 && col+i-1 < len(input[0]); i++ {
		for row := 0; row < len(input); row++ {
			leftCol := col - i
			rightCol := col + i - 1
			if input[row][leftCol] != input[row][rightCol] {
				return false
			}
		}
	}
	return true
}

func findReflect(input [][]rune) []int {
	var reflects []int
	for col := 0; col < len(input[0])-1; col++ {
		if testReflectCol(input, col) {
			reflects = append(reflects, col+1)
		}
	}
	return reflects
}

func q1(input [][][]rune) int {
	var total int
	for _, grid := range input {
		colsReflect := findReflect(grid)

		for _, s := range colsReflect {
			total += s
		}

		gridT := transpose(grid)
		rowsReflect := findReflect(gridT)
		for _, s := range rowsReflect {
			total += 100 * s
		}

	}
	return total
}

func subtractSlice(orig []int, new []int) []int {
	var r []int
	for _, c := range new {
		if !slices.Contains(orig, c) {
			r = append(r, c)
		}
	}
	return r
}
func smudgeMirror(grid [][]rune) ([]int, []int) {
	colReflectOrig := findReflect(grid)
	gridT := transpose(grid)
	rowReflectOrig := findReflect(gridT)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '.' {
				grid[row][col] = '#'
			} else {
				grid[row][col] = '.'
			}

			colReflectNew := findReflect(grid)
			newGridT := transpose(grid)
			rowReflectNew := findReflect(newGridT)

			sRow := subtractSlice(rowReflectOrig, rowReflectNew)
			sCol := subtractSlice(colReflectOrig, colReflectNew)

			if len(sRow) > 0 || len(sCol) > 0 {
				return sRow, sCol
			}

			if grid[row][col] == '.' {
				grid[row][col] = '#'
			} else {
				grid[row][col] = '.'
			}

		}
	}
	return []int{}, []int{}
}

func q2(input [][][]rune) int {
	total := 0
	for _, grid := range input {
		rowReflect, colReflect := smudgeMirror(grid)

		var gridScore int

		if len(colReflect) != 0 {
			gridScore += colReflect[0]
		}
		if len(rowReflect) != 0 {
			gridScore += 100 * rowReflect[0]
		}
		total += gridScore
	}
	return total
}
func main() {
	l := parse("dec13/data/input.txt")
	println(q1(l))

	l = parse("dec13/data/input.txt")
	println(q2(l))
}
