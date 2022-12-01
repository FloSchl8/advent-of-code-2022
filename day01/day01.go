package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	var elvesCalories []int
	tmp := 0
	for _, calory := range splittedCalories {
		if calory != "" {
			i, _ := strconv.Atoi(calory)
			tmp += i
		} else {
			elvesCalories = append(elvesCalories, tmp)
			tmp = 0
		}
	}

	max := 0
	sort.Ints(elvesCalories)
	for _, i := range elvesCalories[len(elvesCalories)-3:] {
		max += i
	}
	return max
}
