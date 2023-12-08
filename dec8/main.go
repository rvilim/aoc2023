package main

import (
	"bufio"
	"os"
)

type Direction struct {
	left  string
	right string
}

func parse(filepath string) ([]rune, map[string]Direction) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineNo := 0
	desertMap := make(map[string]Direction)
	scanner := bufio.NewScanner(file)
	var steps []rune

	for scanner.Scan() {
		line := scanner.Text()

		if lineNo == 0 {
			steps = []rune(line)
		} else if lineNo >= 2 {
			l := []rune(line)
			pos := string(l[:3])
			left := string(l[7:10])
			right := string(l[12:15])
			desertMap[pos] = Direction{left: left, right: right}
		}
		lineNo += 1
	}
	return steps, desertMap

}

func makeStep(currentLoc string, stepNum int, steps []rune, desertMap map[string]Direction) string {
	currentStep := steps[stepNum%len(steps)]

	if currentStep == rune('L') {
		currentLoc = desertMap[currentLoc].left
	} else if currentStep == rune('R') {
		currentLoc = desertMap[currentLoc].right
	}

	return currentLoc

}

func loopLength(currentLoc string, steps []rune, desertMap map[string]Direction) int {
	var ends []int
	stepNum := 0
	for {
		if []rune(currentLoc)[2] == 'Z' {
			ends = append(ends, stepNum)
		}
		if len(ends) == 3 {
			offset := (ends[2] - ends[1]) - (ends[1] - ends[0])
			period := ends[2] - ends[1]

			// The input _seems_ to be easy, you could imagine a universe where the loops are "got into"
			// at differing step nums, so we'd have to add an offset.

			// Also I'm not testing for this but you could also imagine a universe where the parity
			// of the steps matters, e.g. we could return to the same _position_ but at a different point in the
			// steps, which means that the step state matters.

			// In my input neither of these things happened
			if offset != 0 {
				println("Error nonzero offset")
			}
			return period
		}
		currentLoc = makeStep(currentLoc, stepNum, steps, desertMap)
		stepNum += 1
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(nums []int) int {
	result := nums[0] * nums[1] / gcd(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		result = lcm([]int{result, nums[i]})
	}

	return result
}
func q2(steps []rune, desertMap map[string]Direction) int {

	var loops []int
	for k := range desertMap {
		if []rune(k)[2] == 'A' {
			loops = append(loops, loopLength(k, steps, desertMap))
		}
	}

	return lcm(loops)
}

func q1(steps []rune, desertMap map[string]Direction) int {
	stepNum := 0
	currentLoc := "AAA"

	for currentLoc != "ZZZ" {
		currentStep := steps[stepNum%len(steps)]
		if currentStep == rune('L') {
			currentLoc = desertMap[currentLoc].left
		} else if currentStep == rune('R') {
			currentLoc = desertMap[currentLoc].right
		}
		stepNum += 1
	}
	return stepNum
}
func main() {
	steps, desertMap := parse("dec8/data/input.txt")
	println(q1(steps, desertMap))

	steps, desertMap = parse("dec8/data/input.txt")
	println(q2(steps, desertMap))

}
