package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	destStart   int64
	sourceStart int64
	length      int64
}

type Mapping struct {
	dest   string
	ranges []Range
}

func parse(filepath string) ([]int64, map[string]Mapping) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var seeds []int64
	mappings := make(map[string]Mapping)

	mappings["seed"] = Mapping{dest: "soil"}
	mappings["soil"] = Mapping{dest: "fertilizer"}
	mappings["fertilizer"] = Mapping{dest: "water"}
	mappings["water"] = Mapping{dest: "light"}
	mappings["light"] = Mapping{dest: "temperature"}
	mappings["temperature"] = Mapping{dest: "humidity"}
	mappings["humidity"] = Mapping{dest: "location"}

	var currentSection string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sections := []string{"seed-to-soil", "soil-to-fertilizer",
			"fertilizer-to-water", "water-to-light", "light-to-temperature",
			"temperature-to-humidity", "humidity-to-location"}

		if strings.Contains(line, "seeds:") {
			for i, strVal := range strings.Fields(line) {
				if i == 0 { // Skip the first one because it's the label
					continue
				}
				val, _ := strconv.ParseInt(strVal, 10, 64)
				seeds = append(seeds, val)
			}

			continue
		} else if strings.Contains(line, ":") {
			// Here we are just looking for a colon to figure out if it's another section header
			for _, section := range sections {
				if strings.Contains(line, section) {
					currentSection = section
				}
			}
			continue
		}

		switch currentSection {
		case "seeds":
			for _, strVal := range strings.Fields(line) {
				val, _ := strconv.ParseInt(strVal, 10, 64)
				seeds = append(seeds, val)
			}
		default:
			values := strings.Fields(line)
			if len(values) == 3 { // If values is not 3 we have a blank line
				destStart, _ := strconv.ParseInt(values[0], 10, 64)
				sourceStart, _ := strconv.ParseInt(values[1], 10, 64)
				length, _ := strconv.ParseInt(values[2], 10, 64)

				m := strings.Split(currentSection, "-to-")
				sourceName := m[0]

				r := mappings[sourceName]
				r.ranges = append(r.ranges, Range{destStart: destStart, sourceStart: sourceStart, length: length})
				mappings[sourceName] = r
			}
		}
	}
	return seeds, mappings
}

func followMapping(source int64, mappingName string, mappings map[string]Mapping) int64 {
	mapping := mappings[mappingName]
	dest := source

	for _, mappingRange := range mappings[mappingName].ranges {
		if source >= mappingRange.sourceStart && source <= mappingRange.sourceStart+mappingRange.length {
			dest = mappingRange.destStart + (source - mappingRange.sourceStart)
			break
		}
	}

	if mapping.dest == "location" {
		return dest
	} else {
		return followMapping(dest, mapping.dest, mappings)
	}
}

func q1(seeds []int64, mappings map[string]Mapping) int64 {
	minSeed := seeds[0]

	for _, seed := range seeds {
		m := followMapping(seed, "seed", mappings)
		if m < minSeed {
			minSeed = m
		}
	}
	return minSeed
}

func q2(seeds []int64, mappings map[string]Mapping) int64 {
	minSeed := seeds[0]

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{} // Mutex to protect minSeed

	for seedStart := 0; seedStart < len(seeds); seedStart += 2 {

		wg.Add(1)

		go func(start int) {
			defer wg.Done()

			iteration := start / 2
			totalIterations := len(seeds) / 2
			nSeeds := seeds[start] + seeds[start+1] - seeds[start]
			fmt.Printf("starting %d / %d, %d seeds\n", iteration, totalIterations, nSeeds)
			for seed := seeds[start]; seed < seeds[start]+seeds[start+1]; seed++ {

				m := followMapping(seed, "seed", mappings) // Assuming followMapping is defined
				mutex.Lock()
				if m < minSeed {
					minSeed = m
				}
				mutex.Unlock()
			}
			fmt.Printf("finished %d / %d, %d seeds\n", iteration, totalIterations, nSeeds)

		}(seedStart)
	}
	wg.Wait()
	return minSeed

}
func main() {
	seeds, mappings := parse("dec5/data/input.txt")
	println("Q1:")
	println(q1(seeds, mappings))
	println(" ")
	println("Q2:")
	println(q2(seeds, mappings))

}
