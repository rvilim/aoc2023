package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Coord struct {
	row int
	col int
}
type Node struct {
	north bool
	south bool
	east  bool
	west  bool
}

func (m Node) print() rune {
	if m.north && m.south {
		return '│'
	} else if m.north && m.west {
		return '┘'
	} else if m.north && m.east {
		return '└'
	} else if m.east && m.west {
		return '─'
	} else if m.south && m.west {
		return '┐'
	} else if m.south && m.east {
		return '┌'
	}
	return 'X'
}

func isGround(n Node) bool {
	return !(n.north || n.south || n.east || n.west)
}

func parse(filepath string) ([][]Node, Coord) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var results [][]Node
	var start Coord
	var lineNo int

	for scanner.Scan() {
		lineText := scanner.Text()
		lineNodes := make([]Node, len(lineText))
		l := []rune(lineText)

		for i, r := range l {
			switch r {
			case '|':
				lineNodes[i].north = true
				lineNodes[i].south = true
			case '-':
				lineNodes[i].east = true
				lineNodes[i].west = true
			case 'L':
				lineNodes[i].north = true
				lineNodes[i].east = true
			case 'J':
				lineNodes[i].north = true
				lineNodes[i].west = true
			case '7':
				lineNodes[i].west = true
				lineNodes[i].south = true
			case 'F':
				lineNodes[i].east = true
				lineNodes[i].south = true
			case 'S':
				start = Coord{row: lineNo, col: i}
			}
		}
		results = append(results, lineNodes)
		lineNo += 1
	}

	startShape := findStartShape(results, start)
	results[start.row][start.col] = startShape

	return results, start
}

func traverse(pipeMap [][]Node, node Coord, visited []Coord, steps int) (int, []Coord) {
	c := Coord{row: node.row, col: node.col + 1}
	if pipeMap[node.row][node.col].east && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Coord{row: node.row, col: node.col - 1}
	if pipeMap[node.row][node.col].west && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Coord{row: node.row + 1, col: node.col}
	if pipeMap[node.row][node.col].south && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}

	c = Coord{row: node.row - 1, col: node.col}
	if pipeMap[node.row][node.col].north && !slices.Contains(visited, c) {
		visited = append(visited, node)
		return traverse(pipeMap, c, visited, steps+1)
	}
	return steps, append(visited, node)
}

func findStartShape(pipeMap [][]Node, start Coord) Node {
	var startNode Node

	neighbourRow := start.row - 1
	if neighbourRow >= 0 && pipeMap[neighbourRow][start.col].south {
		startNode.north = true
	}

	neighbourRow = start.row + 1
	if neighbourRow < len(pipeMap) && pipeMap[neighbourRow][start.col].north {
		startNode.south = true
	}

	neighbourCol := start.col - 1
	if neighbourCol >= 0 && pipeMap[start.row][neighbourCol].east {
		startNode.west = true
	}

	neighbourCol = start.col + 1
	if neighbourCol < len(pipeMap[0]) && pipeMap[start.row][neighbourCol].west {
		startNode.east = true
	}

	return startNode

}

func q1(pipeMap [][]Node, start Coord) int {
	var visited []Coord
	pathLen, _ := traverse(pipeMap, start, visited, 0)
	return (pathLen + 1) / 2

}

func isInside(pipeMap [][]Node, loopPoints []Coord, point Coord) bool {
	var nCrossings int

	for _, loopPoint := range loopPoints {
		if loopPoint.row == point.row && loopPoint.col > point.col &&
			(pipeMap[loopPoint.row][loopPoint.col].print() == '│' ||
				pipeMap[loopPoint.row][loopPoint.col].print() == '┐' ||
				pipeMap[loopPoint.row][loopPoint.col].print() == '┌') {
			nCrossings += 1

		}
	}
	return nCrossings%2 == 1
}

func q2(pipeMap [][]Node, start Coord) int {
	var visited []Coord
	var insides []Coord
	_, loopPoints := traverse(pipeMap, start, visited, 0)

	for row := 0; row < len(pipeMap); row++ {
		for col := 0; col < len(pipeMap[0]); col++ {
			point := Coord{row, col}
			if !slices.Contains(loopPoints, point) && isInside(pipeMap, loopPoints, point) {
				insides = append(insides, Coord{row, col})
			}
		}
	}
	printMap(pipeMap, insides)
	return len(insides)
}

func printMap(pipeMap [][]Node, insides []Coord) {
	for rowNum, row := range pipeMap {
		for colNum, c := range row {
			if slices.Contains(insides, Coord{rowNum, colNum}) {
				fmt.Print("I")
			} else {
				fmt.Print(string(c.print()))
			}

		}
		fmt.Println("")
	}
}
func main() {
	pipeMap, start := parse("dec10/data/input.txt")
	println(q1(pipeMap, start))

	pipeMap, start = parse("dec10/data/input.txt")
	println(q2(pipeMap, start))
}
