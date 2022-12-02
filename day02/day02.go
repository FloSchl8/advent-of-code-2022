package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("My score:", getScore(string(input)))
}

func getScore(input string) int {
	rounds := strings.Split(input, "\n")

	score := 0
	for _, round := range rounds {
		if round != "" {
			players := strings.Split(round, " ")
			elf, expectedResult := players[0], players[1]
			switch elf {
			case "A": // rock
				switch expectedResult {
				case "X": // loose
					score += 3 // sciccor 0+3
				case "Y": // draw
					score += 4 // rock 1+3
				case "Z": // win
					score += 8 // paper 2+6
				}
			case "B": // paper
				switch expectedResult {
				case "X":
					score += 1
				case "Y":
					score += 5
				case "Z":
					score += 9
				}
			case "C": // scissor
				switch expectedResult {
				case "X":
					score += 2
				case "Y":
					score += 6
				case "Z":
					score += 7
				}
			}
		}
	}

	return score
}
