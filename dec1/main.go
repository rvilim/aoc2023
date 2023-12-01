package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func get_calibration_q1(line string) int {
	var first, last rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			if first == 0 {
				first = c
				last = c
			} else {
				last = c
			}
		}
	}
	calibration, _ := strconv.Atoi(string(first) + string(last))
	return calibration
}

func get_calibration_q2(line string) int {
	searches := []struct {
		name string
		val  string
	}{{name: "one", val: "1"},
		{name: "two", val: "2"},
		{name: "three", val: "3"},
		{name: "four", val: "4"},
		{name: "five", val: "5"},
		{name: "six", val: "6"},
		{name: "seven", val: "7"},
		{name: "eight", val: "8"},
		{name: "nine", val: "9"},
		{name: "1", val: "1"},
		{name: "2", val: "2"},
		{name: "3", val: "3"},
		{name: "4", val: "4"},
		{name: "5", val: "5"},
		{name: "6", val: "6"},
		{name: "7", val: "7"},
		{name: "8", val: "8"},
		{name: "9", val: "9"}}

	var first, last string
	for i := 0; i < len(line); i++ {
		for _, search := range searches {
			if strings.HasPrefix(line[i:], search.name) {
				if first == "" {
					first = search.val
					last = search.val
				} else {
					last = search.val
				}
			}
		}
	}

	calibration, _ := strconv.Atoi(string(first) + string(last))
	return calibration
}

func q1(filepath string) int {
	total := 0

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += get_calibration_q1(scanner.Text())
	}
	return total
}

func q2(filepath string) int {
	total := 0

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total += get_calibration_q2(line)
	}
	return total
}

func main() {
	println(q1("dec1/data/input.txt"))
	println(q2("dec1/data/input.txt"))
}
