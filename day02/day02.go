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
			elf, own := players[0], players[1]
			switch elf {
			case "A":
				switch own {
				case "X":
					score += 4
				case "Y":
					score += 8
				case "Z":
					score += 3
				}
			case "B":
				switch own {
				case "X":
					score += 1
				case "Y":
					score += 5
				case "Z":
					score += 9
				}
			case "C":
				switch own {
				case "X":
					score += 7
				case "Y":
					score += 2
				case "Z":
					score += 6
				}
			}
		}
	}

	return score
}
