package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	numOfCrates int
	pickup      int
	dropAt      int
}

func main() {
	input, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Top crates: ", stackCrates(string(input)))
}

func stackCrates(input string) (topOnes string) {
	lines := strings.Split(input, "\n")

	instructionsStartAfterLine := findSeperatorLine(lines)

	crates := getNumberOfCrates(lines[instructionsStartAfterLine-1])

	stacks := make([][]string, crates)

	for i := 0; i < crates; i++ {
		stacks[i] = make([]string, instructionsStartAfterLine-1)
		for j := 0; j < instructionsStartAfterLine-1; j++ {
			line := lines[j]
			if crate := string([]rune(line)[i*4+1]); crate != " " {
				stacks[i][j] = crate
			}
		}
	}

	fmt.Println(stacks)

	operations := make([]operation, 0)

	for i := instructionsStartAfterLine + 1; i < len(lines); i++ {
		if lines[i] != "" {
			operations = append(operations, getOperation(lines[i]))
		}
	}

	for k, o := range operations {
		pickupStack := stacks[o.pickup-1]
		var movedStack []string
		var newStack []string
		for j := 0; j < len(pickupStack); j++ {
			if pickupStack[j] != "" {
				movedStack, newStack = pickupStack[:j+o.numOfCrates], pickupStack[j+o.numOfCrates:]
				break
			}
		}
		stacks[o.pickup-1] = newStack

		dropAtStack := stacks[o.dropAt-1]
		if len(dropAtStack) == 0 {
			dropAtStack = append(dropAtStack, movedStack...)
		} else {
			for j := 0; j < len(dropAtStack); j++ {
				if dropAtStack[j] != "" {
					firstPart := dropAtStack[:j]
					secondPart := append(movedStack, dropAtStack[j:]...)
					dropAtStack = append(firstPart, secondPart...)
					break
				}
			}
		}
		stacks[o.dropAt-1] = dropAtStack

		fmt.Println("operation: ", k, "stacks: ", stacks)

	}
	result := ""
	for _, stack := range stacks {
		for i := 0; i < len(stack); i++ {
			if stack[i] != "" {
				result += stack[i]
				break
			}
		}
	}

	return result
}

func findSeperatorLine(lines []string) (instructionsStartAfterLine int) {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	return 0
}

func getNumberOfCrates(input string) int {
	return len(strings.Split(input, "   "))
}

func getOperation(operationInput string) operation {
	split := strings.Split(operationInput, " ")

	numOfCrates, _ := strconv.Atoi(split[1])
	pickup, _ := strconv.Atoi(split[3])
	dropAt, _ := strconv.Atoi(split[5])

	return operation{
		numOfCrates: numOfCrates,
		pickup:      pickup,
		dropAt:      dropAt,
	}
}
