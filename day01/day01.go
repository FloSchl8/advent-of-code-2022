package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Most calories:", sumCaolries(string(input)))
}

func sumCaolries(calories string) int {

	splittedCalories := strings.Split(calories, "\n")

	max := 0
	tmp := 0
	for _, calory := range splittedCalories {

		if calory != "" {
			i, _ := strconv.Atoi(calory)
			tmp += i
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}

	return max
}
