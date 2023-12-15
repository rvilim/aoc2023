package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
	"strconv"
	"strings"
)

func parse(filepath string) []string {

	fileContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(fileContent), ",")
}

func hash(input string) int {
	var current int

	for _, c := range input {
		current += int(c)
		current *= 17
		current = current % 256
	}
	return current
}

type lens struct {
	label       string
	focalLength int
}
type op struct {
	opType rune
	lens
}

func parseOps(input []string) []op {
	var ops []op
	for _, s := range input {
		if slices.Contains([]rune(s), '=') {
			a := strings.Split(s, "=")
			num, _ := strconv.Atoi(a[1])
			newOp := op{opType: '=', lens: lens{label: a[0], focalLength: num}}
			ops = append(ops, newOp)
		} else {
			newOp := op{opType: '-', lens: lens{label: s[:len(s)-1]}}
			ops = append(ops, newOp)
		}
	}
	return ops
}
func q1(input []string) int {
	var total int
	for _, i := range input {
		total += hash(i)
	}
	return total
}

func q2(ops []op) int {
	var boxes [256][]lens

	for _, currentOp := range ops {
		currentBoxIndex := hash(currentOp.label)
		// currentBox := boxes[currentBoxIndex]
		if currentOp.opType == '=' {
			found := false
			// If there is already a lens in the box with the same label, replace
			// the old lens with the new lens: remove the old lens and put the new
			// lens in its place, not moving any other lenses in the box.

			for currentLensIndex, currentLens := range boxes[currentBoxIndex] {

				if currentLens.label == currentOp.label {
					boxes[currentBoxIndex][currentLensIndex].focalLength = currentOp.focalLength
					found = true
				}
			}

			// If there is not already a lens in the box with the same label, add the
			// lens to the box immediately behind any lenses already in the box.
			// Don't move any of the other lenses when you do this. If there aren't
			// any lenses in the box, the new lens goes all the way to the front of
			// the box.
			if !found {
				boxes[currentBoxIndex] = append(boxes[currentBoxIndex], currentOp.lens)
			}
		} else {
			var newLenses []lens
			for _, currentLens := range boxes[currentBoxIndex] {
				if currentLens.label != currentOp.label {
					newLenses = append(newLenses, currentLens)
				}
			}
			boxes[currentBoxIndex] = newLenses
		}
	}

	return score2(boxes)
}

func printBoxes(boxes [256][]lens) {
	for boxNum, box := range boxes {
		if len(box) > 0 {
			fmt.Println("Box ", boxNum)
			fmt.Println(box)
		}
	}
}
func score2(boxes [256][]lens) int {
	total := 0

	// Add up the focusing power of all of the lenses. The focusing power of a single lens is the result of multiplying together:

	// One plus the box number of the lens in question.
	// The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
	// The focal length of the lens.

	for boxNumber, box := range boxes {
		for slotNumber, slot := range box {
			a := (1 + boxNumber) * (1 + slotNumber) * slot.focalLength
			total += a
		}
	}
	return total
}

func main() {
	l := parse("dec15/data/input.txt")

	println(q1(l))

	m := parseOps(l)
	println(q2(m))
}
