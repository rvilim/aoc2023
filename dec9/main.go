package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var results [][]int
	for scanner.Scan() {
		splits := strings.Fields(scanner.Text())
		currentLine := make([]int, 0, len(splits))

		for _, r := range splits {
			n, _ := strconv.Atoi(r)
			currentLine = append(currentLine, n)
		}
		results = append(results, currentLine)

	}
	return results
}

func seqDiff(seq []int) []int {
	var s []int
	for i := 0; i < len(seq)-1; i++ {
		s = append(s, seq[i+1]-seq[i])
	}
	return s
}

func allSame(a []int) bool {
	for _, v := range a {
		if v != a[0] {
			return false
		}
	}
	return true
}

func predictSequence(seq []int, predictFirst bool) int {

	if allSame(seq) {
		return seq[0]
	}

	if predictFirst {
		slices.Reverse(seq)
	}

	diffs := [][]int{seq}

	for !allSame(diffs[len(diffs)-1]) {
		diffs = append(diffs, seqDiff(diffs[len(diffs)-1]))
	}

	for i := len(diffs) - 2; i >= 0; i-- {
		extraDiff := diffs[i+1][len(diffs[i+1])-1]
		diffs[i] = append(diffs[i], diffs[i][len(diffs[i])-1]+extraDiff)
	}
	return diffs[0][len(diffs[0])-1]
}

func q1(seqs [][]int) int {
	s := 0
	for _, seq := range seqs {
		s += predictSequence(seq, false)
	}
	return s
}

func q2(seqs [][]int) int {
	s := 0
	for _, seq := range seqs {
		s += predictSequence(seq, true)
	}
	return s
}

func main() {
	seqs := parse("dec9/data/input.txt")
	println(q1(seqs))

	seqs = parse("dec9/data/input.txt")
	println(q2(seqs))

}
