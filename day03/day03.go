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
	fmt.Println("Sum of priorities part 1: ", sortBackpackItems(string(input)))
}

func sortBackpackItems(input string) int {
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

func getCharPriorityFromAsciiIndex(i int32) int {
	if i >= 65 && i <= 90 {
		return int(i%32 + 26)
	} else {
		return int(i % 32)
	}
}
