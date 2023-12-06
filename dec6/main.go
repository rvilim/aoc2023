package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func raceDistance(holdTime int, race Race) int {
	return holdTime * (race.time - holdTime)
}
func q(races []Race) int {
	total := 1
	for _, race := range races {
		var records int

		for holdTime := 0; holdTime <= race.time; holdTime++ {
			if raceDistance(holdTime, race) > race.distance {
				records += 1
			}
		}
		total *= records
	}
	return total
}

func parse1(filepath string) []Race {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var times, distances []string
	var races []Race

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		f := strings.Fields(line)

		if f[0] == "Time:" {
			times = f[1:]
		} else if f[0] == "Distance:" {
			distances = f[1:]
		}

	}

	for i, time := range times {
		t, _ := strconv.Atoi(time)
		d, _ := strconv.Atoi(distances[i])

		races = append(races, Race{
			time:     t,
			distance: d,
		})
	}
	return races

}

func parse2(filepath string) []Race {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var times, distances []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		f := strings.Fields(line)

		if f[0] == "Time:" {
			times = f[1:]
		} else if f[0] == "Distance:" {
			distances = f[1:]
		}

	}
	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))

	return []Race{{
		time:     t,
		distance: d,
	}}
}

func main() {
	races1 := parse1("dec6/data/input.txt")
	println(q(races1))

	races2 := parse2("dec6/data/input.txt")
	println(q(races2))

}
