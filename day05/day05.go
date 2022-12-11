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

	for i, stack := range stacks {
		var tmp []string
		copy(tmp, stack)
		for _, s := range stack {
			if s != "" {
				tmp = append(tmp, s)
			}
		}
		stacks[i] = tmp
	}

	fmt.Println(stacks)

	operations := make([]operation, 0)

	for i := instructionsStartAfterLine + 1; i < len(lines); i++ {
		if lines[i] != "" {
			operations = append(operations, getOperation(lines[i]))
		}
	}

	for k, o := range operations {
		fmt.Println(o)

		movingStack := stacks[o.pickup-1][:o.numOfCrates]

		for i, j := 0, len(movingStack)-1; i < j; i, j = i+1, j-1 {
			movingStack[i], movingStack[j] = movingStack[j], movingStack[i]
		}

		fmt.Println(movingStack)

		for _, crate := range movingStack {
			stacks[o.pickup-1] = stacks[o.pickup-1][1:]
			stacks[o.dropAt-1] = append([]string{crate}, stacks[o.dropAt-1]...)
		}

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
