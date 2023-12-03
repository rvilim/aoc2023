package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type NumberRange struct {
	row   int
	start int
	end   int
	value int
}

func parse(filepath string) []string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		println("Error reading file")
	}

	return strings.Split(string(content), "\n")
}

func parseNumbers(schematic []string, numbers *[]NumberRange) {
	// In all my lines, find any numbers and put them into a slice with start ends and values
	r, _ := regexp.Compile(`\d+`)

	for row, line := range schematic {

		for _, charRange := range r.FindAllStringSubmatchIndex(line, -1) {
			a, _ := strconv.Atoi(line[charRange[0]:charRange[1]])
			*numbers = append(*numbers, NumberRange{row, charRange[0], charRange[1], a})
		}
	}

}

func numberValue(number NumberRange, schematic []string) int {
	// This function returns the value of a number if there are any symbols adjacent to that number

	// For each number range we are looking for adjacent symbols. We start with the row above and below
	rows := []int{number.row - 1, number.row + 1}
	for _, row := range rows {
		for col := number.start - 1; col <= number.end; col++ {
			isInRange := row > 0 && row < len(schematic) && col >= 0 && col < len(schematic[0])
			if isInRange && schematic[row][col] != '.' && !unicode.IsDigit(rune(schematic[row][col])) {
				return number.value
			}
		}
	}

	// We then just need to check the ends on the same row
	cols := []int{number.start - 1, number.end}
	for _, col := range cols {
		isInRange := col >= 0 && col < len(schematic[0])
		if isInRange && schematic[number.row][col] != '.' && !unicode.IsDigit(rune(schematic[number.row][col])) {
			return number.value
		}
	}
	return 0
}

func q1(schematic []string) int {
	var numbers []NumberRange
	total := 0

	parseNumbers(schematic, &numbers)

	for _, number := range numbers {
		total += numberValue(number, schematic)
	}

	return total
}

func q2(schematic []string) int {
	var numbers []NumberRange
	total := 0

	parseNumbers(schematic, &numbers)

	// We do this the opposeit way as q1 instead of looping over all the numbers we lop over all
	// the symbols, check if it's a "*", then find all the numbers that border it (by looping over the numbers)
	// as we find an adjacent number we append that to a slice, then at the end check if the length of that slice
	// is 2, it it is we add the product of those two numbers to our running total.
	for row, line := range schematic {
		for col, char := range line {
			if char == '*' {
				var adjacentNumbers []NumberRange
				for _, number := range numbers {
					if number.start <= col+1 && number.end >= col && (number.row == row-1 || number.row == row+1) {
						adjacentNumbers = append(adjacentNumbers, number)
					}
					if row == number.row && (number.end == col || number.start == col+1) {
						adjacentNumbers = append(adjacentNumbers, number)
					}
				}
				if len(adjacentNumbers) == 2 {
					total += adjacentNumbers[0].value * adjacentNumbers[1].value
				}
			}

		}
	}
	return total
}
func main() {
	schematic := parse("dec3/data/input.txt")

	println(q1(schematic))
	println(q2(schematic))
}
