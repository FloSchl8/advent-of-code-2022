package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1", "Signal after:", part01(string(input)))
}

func part01(input string) int {

	for i := 0; i < len(input); i++ {
		set := make(map[string]interface{})
		seq := input[i : i+4]
		for _, s := range seq {
			set[string(s)] = ""
		}
		if len(set) == 4 {
			return i + 4
		}
	}
	return 0
}
