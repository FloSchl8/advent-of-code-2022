package main

import (
	"fmt"
	"sort"
)

type monkey struct {
	items      []int
	operation  func(old int) int
	test       func(worrylevel int) bool
	trueThrow  int
	falseThrow int
}

func main() {

	monkeys := []*monkey{{
		items: []int{96, 60, 68, 91, 83, 57, 85},
		operation: func(old int) int {
			return old * 2
		},
		test: func(worrylevel int) bool {
			return worrylevel%17 == 0
		},
		trueThrow:  2,
		falseThrow: 5,
	}, {
		items: []int{75, 78, 68, 81, 73, 99},
		operation: func(old int) int {
			return old + 3
		},
		test: func(worrylevel int) bool {
			return worrylevel%13 == 0
		},
		trueThrow:  7,
		falseThrow: 4,
	}, {
		items: []int{69, 86, 67, 55, 96, 69, 94, 85},
		operation: func(old int) int {
			return old + 6
		},
		test: func(worrylevel int) bool {
			return worrylevel%19 == 0
		},
		trueThrow:  6,
		falseThrow: 5,
	}, {
		items: []int{88, 75, 74, 98, 80},
		operation: func(old int) int {
			return old + 5
		},
		test: func(worrylevel int) bool {
			return worrylevel%7 == 0
		},
		trueThrow:  7,
		falseThrow: 1,
	}, {
		items: []int{82},
		operation: func(old int) int {
			return old + 8
		},
		test: func(worrylevel int) bool {
			return worrylevel%11 == 0
		},
		trueThrow:  0,
		falseThrow: 2,
	}, {
		items: []int{72, 92, 92},
		operation: func(old int) int {
			return old * 5
		},
		test: func(worrylevel int) bool {
			return worrylevel%3 == 0
		},
		trueThrow:  6,
		falseThrow: 3,
	}, {
		items: []int{74, 61},
		operation: func(old int) int {
			return old * old
		},
		test: func(worrylevel int) bool {
			return worrylevel%2 == 0
		},
		trueThrow:  3,
		falseThrow: 1,
	}, {
		items: []int{76, 86, 83, 55},
		operation: func(old int) int {
			return old + 4
		},
		test: func(worrylevel int) bool {
			return worrylevel%5 == 0
		},
		trueThrow:  4,
		falseThrow: 0,
	},
	}

	fmt.Println("Part 1", "Monkey business", part01(monkeys))
}

func part01(monkeys []*monkey) int {

	monkeyCounter := map[int]int{}

	for i := 0; i < len(monkeys); i++ {
		monkeyCounter[i] = 0
	}

	for i := 0; i < 20; i++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				newLevel := m.operation(item) / 3
				monkeyCounter[i]++
				if m.test(newLevel) {
					monkeys[m.trueThrow].items = append(monkeys[m.trueThrow].items, newLevel)
				} else {
					monkeys[m.falseThrow].items = append(monkeys[m.falseThrow].items, newLevel)
				}
			}
			m.items = nil
		}
	}

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
