package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items      []int
	operation  func(old int) int
	testMod    int
	trueThrow  int
	falseThrow int
}

func main() {

	monkeys := []*monkey{{
		items: []int{96, 60, 68, 91, 83, 57, 85},
		operation: func(old int) int {
			return old * 2
		},
		testMod:    17,
		trueThrow:  2,
		falseThrow: 5,
	}, {
		items: []int{75, 78, 68, 81, 73, 99},
		operation: func(old int) int {
			return old + 3
		},
		testMod:    13,
		trueThrow:  7,
		falseThrow: 4,
	}, {
		items: []int{69, 86, 67, 55, 96, 69, 94, 85},
		operation: func(old int) int {
			return old + 6
		},
		testMod:    19,
		trueThrow:  6,
		falseThrow: 5,
	}, {
		items: []int{88, 75, 74, 98, 80},
		operation: func(old int) int {
			return old + 5
		},
		testMod:    7,
		trueThrow:  7,
		falseThrow: 1,
	}, {
		items: []int{82},
		operation: func(old int) int {
			return old + 8
		},
		testMod:    11,
		trueThrow:  0,
		falseThrow: 2,
	}, {
		items: []int{72, 92, 92},
		operation: func(old int) int {
			return old * 5
		},
		testMod:    3,
		trueThrow:  6,
		falseThrow: 3,
	}, {
		items: []int{74, 61},
		operation: func(old int) int {
			return old * old
		},
		testMod:    2,
		trueThrow:  3,
		falseThrow: 1,
	}, {
		items: []int{76, 86, 83, 55},
		operation: func(old int) int {
			return old + 4
		},
		testMod:    5,
		trueThrow:  4,
		falseThrow: 0,
	},
	}

	fmt.Println("Part 1", "Monkey business", part01(monkeys))
	// since we're using pointers our initial values need to be restored
	monkeys = []*monkey{{
		items: []int{96, 60, 68, 91, 83, 57, 85},
		operation: func(old int) int {
			return old * 2
		},
		testMod:    17,
		trueThrow:  2,
		falseThrow: 5,
	}, {
		items: []int{75, 78, 68, 81, 73, 99},
		operation: func(old int) int {
			return old + 3
		},
		testMod:    13,
		trueThrow:  7,
		falseThrow: 4,
	}, {
		items: []int{69, 86, 67, 55, 96, 69, 94, 85},
		operation: func(old int) int {
			return old + 6
		},
		testMod:    19,
		trueThrow:  6,
		falseThrow: 5,
	}, {
		items: []int{88, 75, 74, 98, 80},
		operation: func(old int) int {
			return old + 5
		},
		testMod:    7,
		trueThrow:  7,
		falseThrow: 1,
	}, {
		items: []int{82},
		operation: func(old int) int {
			return old + 8
		},
		testMod:    11,
		trueThrow:  0,
		falseThrow: 2,
	}, {
		items: []int{72, 92, 92},
		operation: func(old int) int {
			return old * 5
		},
		testMod:    3,
		trueThrow:  6,
		falseThrow: 3,
	}, {
		items: []int{74, 61},
		operation: func(old int) int {
			return old * old
		},
		testMod:    2,
		trueThrow:  3,
		falseThrow: 1,
	}, {
		items: []int{76, 86, 83, 55},
		operation: func(old int) int {
			return old + 4
		},
		testMod:    5,
		trueThrow:  4,
		falseThrow: 0,
	},
	}
	fmt.Println("Part 2", "Monkey business", part02(monkeys))

}

func part01(monkeys []*monkey) int {

	monkeyCounter := playMonkeyBuisness(monkeys, 20, 0)

	result := 1
	ints := []int{}
	for _, i := range monkeyCounter {
		ints = append(ints, i)
	}
	sort.Ints(ints)

	for _, i := range ints[len(ints)-2:] {
		result *= i
	}

	return result
}

func part02(monkeys []*monkey) int {

	commonDiv := 1

	for _, m := range monkeys {
		commonDiv *= m.testMod
	}

	monkeyCounter := playMonkeyBuisness(monkeys, 10000, commonDiv)

	result := 1
	ints := []int{}
	for _, i := range monkeyCounter {
		ints = append(ints, i)
	}
	sort.Ints(ints)

	for _, i := range ints[len(ints)-2:] {
		result *= i
	}

	return result
}

func playMonkeyBuisness(monkeys []*monkey, rounds, commonDiv int) map[int]int {
	monkeyCounter := map[int]int{}

	for i := 0; i < len(monkeys); i++ {
		monkeyCounter[i] = 0
	}

	for j := 0; j < rounds; j++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				newLevel := 0
				if commonDiv == 0 {
					newLevel = m.operation(item) / 3
				} else {
					newLevel = m.operation(item) % commonDiv
				}
				monkeyCounter[i]++
				if newLevel%m.testMod == 0 {
					monkeys[m.trueThrow].items = append(monkeys[m.trueThrow].items, newLevel)
				} else {
					monkeys[m.falseThrow].items = append(monkeys[m.falseThrow].items, newLevel)
				}
			}
			m.items = nil
		}
		if j%1000 == 0 {
			fmt.Println("Round", j, monkeyCounter)
		}
	}
	return monkeyCounter
}
