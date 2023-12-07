package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func parse(filepath string) []Hand {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		f := strings.Fields(line)

		hand := f[0]
		bid, _ := strconv.Atoi(f[1])

		hands = append(hands, Hand{cards: hand, bid: bid})

	}

	return hands
}

func cardToJackHigh(card rune) int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return 9
	case 'T':
		return 8
	default:
		idx, _ := strconv.Atoi(string(card))
		return idx - 2
	}

}

func cardToJackLow(card rune) int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'T':
		return 9
	case 'J':
		return 0
	default:
		idx, _ := strconv.Atoi(string(card))
		return idx - 1
	}

}

func cardsSignature(cards string) []int {
	// Returns a slice of the counts of cards in a hand, sorted most occuring to least occuring
	s := make([]int, 13)

	for _, card := range cards {
		s[cardToJackHigh(card)] += 1
	}

	slices.Sort(s)

	// Reverse it to make the highest occurance first, this will just be helpful later
	slices.Reverse(s)
	return s
}

func scoreHandType(cards string) int {
	// The logic here is that I'm going to first count the number of distinct cards in each hand and sort that
	// which will give me a hand signature e.g.

	// Five of a kind: [0, ... , 0, 5]
	// Four of a kind: [0, ... , 1, 4]
	// Full House    : [0, ... , 2, 3]

	// I'll just return an integer that increases with better hands

	s := cardsSignature(cards)

	if s[0] == 5 { // five of a kind
		return 6
	} else if s[0] == 4 && s[1] == 1 { // four of a kind
		return 5
	} else if s[0] == 3 && s[1] == 2 { //full house
		return 4
	} else if s[0] == 3 && s[1] == 1 && s[2] == 1 { // three of a kind, the last two checks are not strictly neccesary
		return 3
	} else if s[0] == 2 && s[1] == 2 && s[2] == 1 { // two pair, the last check is not strictly neccesary
		return 2
	} else if s[0] == 2 && s[1] == 1 && s[2] == 1 && s[3] == 1 { // one pair, the last three checks are not neccesary
		return 1
	} else if s[0] == 1 && s[1] == 1 && s[2] == 1 && s[3] == 1 && s[4] == 1 { // high card
		return 0
	} else {
		println("Error")
		return -1
	}
}

func q1(hands []Hand) int {
	slices.SortFunc(hands, cmpQ1)
	winnings := 0

	for i, hand := range hands {
		rank := i + 1
		winnings += rank * hand.bid
	}

	return winnings
}

func findStrongestHandType(cards string, pos int) int {
	strongestHandTypeScore := 0
	if pos == 5 {
		return scoreHandType(cards)
	}

	replaceCards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	if cards[pos] != 'J' {
		return findStrongestHandType(cards, pos+1)
	} else {
		for _, replaceCard := range replaceCards {
			s := cards[:pos] + string(replaceCard) + cards[pos+1:]
			handTypeScore := findStrongestHandType(s, pos+1)

			if handTypeScore > strongestHandTypeScore {
				strongestHandTypeScore = handTypeScore
			}
		}
	}
	return strongestHandTypeScore
}

func cmpQ1(hand1 Hand, hand2 Hand) int {
	handTypeScore1 := scoreHandType(hand1.cards)
	handTypeScore2 := scoreHandType(hand2.cards)

	if handTypeScore1 > handTypeScore2 {
		return 1
	} else if handTypeScore1 < handTypeScore2 {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			val1 := cardToJackHigh([]rune(hand1.cards)[i])
			val2 := cardToJackHigh([]rune(hand2.cards)[i])
			if val1 > val2 {
				return 1
			} else if val1 < val2 {
				return -1
			}
		}
	}
	return 0
}

func cmpQ2(hand1 Hand, hand2 Hand) int {
	handTypeScore1 := findStrongestHandType(hand1.cards, 0)
	handTypeScore2 := findStrongestHandType(hand2.cards, 0)

	if handTypeScore1 > handTypeScore2 {
		return 1
	} else if handTypeScore1 < handTypeScore2 {
		return -1
	} else {
		for i := 0; i < 5; i++ {
			val1 := cardToJackLow([]rune(hand1.cards)[i])
			val2 := cardToJackLow([]rune(hand2.cards)[i])
			if val1 > val2 {
				return 1
			} else if val1 < val2 {
				return -1
			}
		}
	}
	return 0
}

func q2(hands []Hand) int {
	slices.SortFunc(hands, cmpQ2)
	winnings := 0

	for i, hand := range hands {
		rank := i + 1
		winnings += rank * hand.bid
	}
	return winnings
}
func main() {
	hands := parse("dec7/data/input.txt")

	println(q1(hands))
	println(q2(hands))

}
