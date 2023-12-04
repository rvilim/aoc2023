package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	have    []int
	winning []int
}

func parse(filepath string) []Card {
	content, err := os.ReadFile(filepath)

	var cards []Card
	if err != nil {
		fmt.Println("Error reading file")
	}
	for _, line := range strings.Split(string(content), "\n") {
		pipePos := strings.Index(line, "|")
		r, _ := regexp.Compile(`\d+`)

		var card Card
		for i, charRange := range r.FindAllStringSubmatchIndex(line, -1) {
			num, _ := strconv.Atoi(line[charRange[0]:charRange[1]])

			if charRange[1] < pipePos && i > 0 {
				card.winning = append(card.winning, num)
			} else if charRange[1] >= pipePos {
				card.have = append(card.have, num)
			}

		}
		cards = append(cards, card)
	}
	return cards
}

func q1(cards []Card) int {
	points := 0

	for _, card := range cards {
		cardPoints := 0
		for _, haveNumber := range card.have {
			if slices.Contains(card.winning, haveNumber) {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}
		points += cardPoints

	}
	return points
}

func q2(cards []Card) int {
	type CardCopy struct {
		card   Card
		copies int
	}
	var cardCopies []CardCopy

	for _, card := range cards {
		cardCopies = append(cardCopies, CardCopy{card: card, copies: 1})
	}

	for cardIndex, card := range cardCopies {
		cardMatches := 0
		for _, haveNumber := range card.card.have {
			if slices.Contains(card.card.winning, haveNumber) {
				cardMatches += 1
			}
		}

		for i := 1; i < cardMatches+1; i++ {
			cardCopies[cardIndex+i].copies += card.copies
		}
	}

	total := 0
	for _, cardCopy := range cardCopies {
		total += cardCopy.copies
	}
	return total
}

func main() {
	cards := parse("dec4/data/input.txt")
	println(q1(cards))
	println(q2(cards))
}
