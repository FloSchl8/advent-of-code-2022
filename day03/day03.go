package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum of priorities part 1: ", part1(string(input)))
	fmt.Println("Sum of priorities part 2: ", part2(string(input)))

}

func part1(input string) int {
	backpacks := strings.Split(input, "\n")

	sum := 0
	for _, backpack := range backpacks {
		if backpack != "" {
			first, second := backpack[:len(backpack)/2], backpack[len(backpack)/2:]
			fmt.Println("first", first, "second", second)
			for _, c := range first {
				if strings.Contains(second, string(c)) {
					sum += getCharPriorityFromAsciiIndex(c)
					break
				}
			}
		}
	}

	return sum
}

func part2(input string) int {
	backpacks := strings.Split(input, "\n")
	sum := 0
	for i := 0; i < len(backpacks); i += 3 {
		for _, c := range backpacks[i] {
			if strings.Contains(backpacks[i+1], string(c)) && strings.Contains(backpacks[i+2], string(c)) {
				sum += getCharPriorityFromAsciiIndex(c)
				break
			}
		}
	}
	return sum
}

func getCharPriorityFromAsciiIndex(i int32) int {
	if i >= 65 && i <= 90 {
		return int(i%32 + 26)
	} else {
		return int(i % 32)
	}
}
