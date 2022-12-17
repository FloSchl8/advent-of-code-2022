package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1", "Visible trees:", getVisibleTrees(string(input)))
}

func getVisibleTrees(input string) int {
	lines := strings.Split(input, "\n")

	outerTrees := getOuterTrees(len(lines[0]))

	innerTrees := getInnerTrees(lines)

	return outerTrees + innerTrees
}

func getInnerTrees(lines []string) int {

	result := 0

	for i, line := range lines {
		for j, tree := range line {
			if line != "" && i > 0 && j > 0 && i < len(line)-1 && j < len(line)-1 {
				treeHeight, _ := strconv.Atoi(string(tree))
				verticalLine := getLineHeights(line)
				horizontalLine := getHorizontalLineHeights(lines, j)

				if treeIsVisible(treeHeight, horizontalLine, verticalLine, i, j) {
					result++
				}
			}
		}
	}
	return result
}

func getHorizontalLineHeights(lines []string, index int) []int {
	tmpLine := ""
	for _, line := range lines {
		if line != "" {
			tmpLine += string(line[index])
		}
	}
	return getLineHeights(tmpLine)
}

func getLineHeights(line string) []int {
	var result = []int{}
	for k := 0; k < len(line); k++ {
		h, _ := strconv.Atoi(string(line[k]))
		result = append(result, h)
	}
	return result
}

func treeIsVisible(treeHeight int, horizontalLine, verticalLine []int, x, y int) bool {

	// horizontal starting at tree index walking 'right' ending
	result := false
	for i := x + 1; i < len(horizontalLine); i++ {
		if treeHeight > horizontalLine[i] {
			result = true
		} else {
			result = false
			break
		}
	}

	if !result {
		// horizontal starting at zero till tree index
		for i := 0; i < x; i++ {
			if treeHeight > horizontalLine[i] {
				result = true
			} else {
				result = false
				break
			}
		}
	}

	if !result {
		// vertical starting at tree index walking 'right' ending
		for i := y + 1; i < len(verticalLine); i++ {
			if treeHeight > verticalLine[i] {
				result = true
			} else {
				result = false
				break
			}
		}
	}
	if !result {
		// vertical starting at zero till tree index
		for i := 0; i < y; i++ {
			if treeHeight > verticalLine[i] {
				result = true
			} else {
				result = false
				break
			}
		}
	}
	return result
}

func getOuterTrees(i int) int {
	return (2 * i) + 2*(i-2)
}
