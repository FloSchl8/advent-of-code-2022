package main

import (
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1", "Sum of cycle * register", part01(string(input)))
	fmt.Println("Part 2", "CRT displays:")
	fmt.Println(part02(string(input)))

}

func part01(input string) int {
	lines := strings.Split(input, "\n")

	register := 1
	cycle := 1

	cycleRegisterMap := map[int]int{}
	// wanted register values
	cycleRegisterMap[20] = -1
	cycleRegisterMap[60] = -1
	cycleRegisterMap[100] = -1
	cycleRegisterMap[140] = -1
	cycleRegisterMap[180] = -1
	cycleRegisterMap[220] = -1

	targetCycles := make([]int, 0, len(cycleRegisterMap))
	for k := range cycleRegisterMap {
		targetCycles = append(targetCycles, k)
	}

	for _, line := range lines {
		programline := strings.Split(line, " ")

		// if we hit a cycle direct, add the value
		// if we get a noop this time and stand and targetCycle-1, we'll get it in the next cycle
		// with addx we handle it in the else branch
		if cycleRegisterMap[cycle] == -1 {
			cycleRegisterMap[cycle] = register
		}

		if len(programline) == 1 { // noop
			cycle++
		} else { // addx
			for _, targetCycle := range targetCycles {
				// we would jump over our target cycle and wouldn't add it
				if cycle+1 == targetCycle {
					cycleRegisterMap[cycle+1] = register
				}
			}
			cycle += 2
			value, _ := strconv.Atoi(programline[1])
			register += value
		}
	}

	result := 0
	for c, r := range cycleRegisterMap {
		result += c * r
	}

	return result
}

func part02(input string) string {
	lines := strings.Split(input, "\n")

	register := 1

	r := ring.New(240)
	crtOutput := ""
	lineIndex := 0
	nextOperationAtCycle := 0
	operationValue := 0

	for i := 0; i < r.Len(); i++ {
		if i%40 == 0 && i > 0 {
			crtOutput += "\n"
		}

		if i == nextOperationAtCycle {
			register += operationValue
			programline := strings.Split(lines[lineIndex], " ")
			lineIndex++
			if len(programline) == 1 {
				nextOperationAtCycle++
				operationValue = 0
			} else { // addx
				nextOperationAtCycle += 2
				operationValue, _ = strconv.Atoi(programline[1])
			}
		}

		if iMod := i % 40; iMod == register-1 || iMod == register || iMod == register+1 {
			crtOutput += "#"
		} else {
			crtOutput += "."
		}
	}

	return crtOutput
}
