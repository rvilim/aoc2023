package main

import (
	"bufio"
	"os"
)

func parse(filepath string) ([][]rune, [][][2]int, [][][2]int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var platform [][]rune

	for scanner.Scan() {
		platform = append(platform, []rune(scanner.Text()))

	}

	var rowBounds [][][2]int
	var colBounds [][][2]int

	for _, row := range platform {
		r := parseLine(row)
		rowBounds = append(rowBounds, r)
	}

	for colNum := 0; colNum < len(platform[0]); colNum++ {
		col := extractColumn(platform, colNum)
		c := parseLine(col)
		colBounds = append(colBounds, c)
	}
	return platform, rowBounds, colBounds
}

func parseLine(line []rune) [][2]int {
	var bounds [][2]int
	inPeriod := false
	start := 0

	for i, char := range line {
		if char != '#' {
			if !inPeriod {
				start = i
				inPeriod = true
			}
		} else {
			if inPeriod {
				bounds = append(bounds, [2]int{start, i - 1})
				inPeriod = false
			}
		}
	}

	// Check if the last character is a period
	if inPeriod {
		bounds = append(bounds, [2]int{start, len(line) - 1})
	}

	return bounds
}

func extractColumn(input [][]rune, colNum int) []rune {
	var a []rune
	for i := 0; i < len(input); i++ {
		a = append(a, input[i][colNum])
	}
	return a
}
func score(input [][]rune) int {
	total := 0
	for row, line := range input {
		for _, c := range line {
			if c == 'O' {
				total += (len(input) - row)
			}
		}
	}
	return total
}

func tiltWest(platform [][]rune, rowBounds [][][2]int) [][]rune {
	for rowNum := 0; rowNum < len(platform); rowNum++ {
		for _, bounds := range rowBounds[rowNum] {
			nO := countRow(platform[rowNum], bounds)
			for i := bounds[0]; i <= bounds[1]; i++ {
				if nO > 0 {
					platform[rowNum][i] = 'O'
					nO -= 1
				} else {
					platform[rowNum][i] = '.'
				}
			}
		}
	}
	return platform
}

func countRow(row []rune, bounds [2]int) int {
	nO := 0
	for i := bounds[0]; i <= bounds[1]; i++ {
		if row[i] == 'O' {
			nO += 1
		}
	}
	return nO
}

func countCol(platform [][]rune, colNum int, bounds [2]int) int {
	nO := 0
	for i := bounds[0]; i <= bounds[1]; i++ {
		if platform[i][colNum] == 'O' {
			nO += 1
		}
	}
	return nO
}

func tiltEast(platform [][]rune, rowBounds [][][2]int) [][]rune {
	for rowNum := 0; rowNum < len(platform); rowNum++ {
		for _, bounds := range rowBounds[rowNum] {
			nO := countRow(platform[rowNum], bounds)
			for i := bounds[1]; i >= bounds[0]; i-- {
				if nO > 0 {
					platform[rowNum][i] = 'O'
					nO -= 1
				} else {
					platform[rowNum][i] = '.'
				}
			}
		}
	}
	return platform
}

func tiltSouth(platform [][]rune, colBounds [][][2]int) [][]rune {
	for colNum := 0; colNum < len(platform[0]); colNum++ {
		for _, bounds := range colBounds[colNum] {
			nO := countCol(platform, colNum, bounds)
			for i := bounds[1]; i >= bounds[0]; i-- {
				if nO > 0 {
					platform[i][colNum] = 'O'
					nO -= 1
				} else {
					platform[i][colNum] = '.'
				}
			}
		}
	}
	return platform
}

func tiltNorth(platform [][]rune, colBounds [][][2]int) [][]rune {
	for colNum := 0; colNum < len(platform[0]); colNum++ {
		for _, bounds := range colBounds[colNum] {
			nO := countCol(platform, colNum, bounds)
			for i := bounds[0]; i <= bounds[1]; i++ {
				if nO > 0 {
					platform[i][colNum] = 'O'
					nO -= 1
				} else {
					platform[i][colNum] = '.'
				}
			}
		}
	}
	return platform
}

func q1(input [][]rune, rowBounds [][][2]int, colBounds [][][2]int) int {
	return score(tiltNorth(input, colBounds))
}

func runesToString(runes [][]rune) string {
	var result string
	for _, row := range runes {
		result += string(row)
	}
	return result
}

func cycle(platform [][]rune, rowBounds [][][2]int, colBounds [][][2]int) [][]rune {
	res := tiltEast(tiltSouth(tiltWest(tiltNorth(platform, colBounds), rowBounds), colBounds), rowBounds)
	return res
}

func q2(platform [][]rune, rowBounds [][][2]int, colBounds [][][2]int) int {
	scoreHistory := []int{0}
	cycleCache := make(map[string]int)
	i := 1
	nIter := 1000000000
	for {
		platform := cycle(platform, rowBounds, colBounds)
		scoreHistory = append(scoreHistory, score(platform))
		_, ok := cycleCache[runesToString(platform)]

		if ok {
			break
		}
		cycleCache[runesToString(platform)] = i
		i += 1
	}
	cycleLength := i - cycleCache[runesToString(platform)]
	cycleStart := cycleCache[runesToString(platform)]

	t := cycleStart + (nIter-cycleStart)%cycleLength
	return scoreHistory[t]
}

func main() {
	l, rowBounds, colBounds := parse("dec14/data/input.txt")

	println(q1(l, rowBounds, colBounds))
	println(q2(l, rowBounds, colBounds))
}
