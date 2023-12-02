package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	GameNumber int
	Rounds     []CubeSet
}

func getGame(line string) Game {
	var g Game

	parts := strings.SplitN(line, ": ", 2)
	gameInfo := strings.Fields(parts[0])
	g.GameNumber, _ = strconv.Atoi(gameInfo[1])

	rounds := strings.Split(parts[1], "; ")

	for _, round := range rounds {
		draw := strings.Split(round, ", ")
		var d CubeSet
		for _, c := range draw {
			s := strings.SplitN(c, " ", 2)
			number, _ := strconv.Atoi(s[0])
			colour := s[1]

			switch colour {
			case "red":
				d.Red = number

			case "blue":
				d.Blue = number

			case "green":
				d.Green = number
			}
		}
		g.Rounds = append(g.Rounds, d)
	}
	return g
}

func Q2GetPower(line string) int {
	c := CubeSet{}
	game := getGame(line)

	for _, round := range game.Rounds {
		if round.Red > c.Red {
			c.Red = round.Red
		}
		if round.Blue > c.Blue {
			c.Blue = round.Blue
		}
		if round.Green > c.Green {
			c.Green = round.Green
		}
	}
	return c.Blue * c.Red * c.Green
}

func Q1GetGamePoints(line string) int {
	c := CubeSet{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	game := getGame(line)

	for _, round := range game.Rounds {

		if round.Blue > c.Blue || round.Red > c.Red || round.Green > c.Green {
			return 0
		}
	}
	return game.GameNumber
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
		total += Q1GetGamePoints(scanner.Text())
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
		total += Q2GetPower(scanner.Text())
	}
	return total
}

func main() {
	println(q1("dec2/data/input.txt"))
	println(q2("dec2/data/input.txt"))
}
