package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section struct {
	lowerEnd int
	upperEnd int
}

func newSectionFromBound(pair string) section {
	lower, upper := splitPairInBounds(pair)
	return section{
		lowerEnd: lower,
		upperEnd: upper,
	}
}

func (s section) isContained(comparedSection section) bool {
	return (s.lowerEnd <= comparedSection.lowerEnd && s.upperEnd >= comparedSection.upperEnd) ||
		(s.lowerEnd >= comparedSection.lowerEnd && s.upperEnd <= comparedSection.upperEnd)
}

func (s section) isOverlapping(comparedSection section) bool {
	return (s.lowerEnd <= comparedSection.lowerEnd && s.lowerEnd >= comparedSection.upperEnd) ||
		(s.upperEnd >= comparedSection.lowerEnd && s.upperEnd <= comparedSection.upperEnd) ||
		(s.lowerEnd >= comparedSection.lowerEnd && s.lowerEnd <= comparedSection.upperEnd) ||
		(s.upperEnd <= comparedSection.lowerEnd && s.upperEnd >= comparedSection.upperEnd)
}

func main() {
	input, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1 - Number of fully conained ranges: ", findFullyContainedRanges(string(input)))
	fmt.Println("Part 2 - Number of fully conained ranges: ", findOverallpingRanges(string(input)))

}

func findFullyContainedRanges(input string) int {
	ranges := strings.Split(input, "\n")

	count := 0

	for _, pair := range ranges {
		if pair != "" {
			splittedRanges := strings.Split(pair, ",")
			if newSectionFromBound(splittedRanges[0]).isContained(newSectionFromBound(splittedRanges[1])) {
				count++
			}
		}
	}
	return count
}

func findOverallpingRanges(input string) int {
	ranges := strings.Split(input, "\n")

	count := 0

	for _, pair := range ranges {
		if pair != "" {
			splittedRanges := strings.Split(pair, ",")
			if newSectionFromBound(splittedRanges[0]).isOverlapping(newSectionFromBound(splittedRanges[1])) || newSectionFromBound(splittedRanges[0]).isContained(newSectionFromBound(splittedRanges[1])) {
				count++
			}
		}
	}
	return count
}

func splitPairInBounds(rangeInput string) (lower int, upper int) {
	borders := strings.Split(rangeInput, "-")
	lower, err := strconv.Atoi(borders[0])
	if err != nil {
		log.Fatal(err)
	}
	upper, err = strconv.Atoi(borders[1])
	if err != nil {
		log.Fatal(err)
	}

	return lower, upper
}
