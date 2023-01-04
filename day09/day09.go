package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type move struct {
	steps     int
	direction string
}

type point struct {
	x, y int
}

func (t *tail) distanceFromHead(head head) float64 {
	return math.Sqrt(math.Pow(float64(head.x-t.x), 2) + math.Pow(float64(head.y-t.y), 2))
}

func (p *point) pointIsEqual(other point) bool {
	return p.x == other.x && p.y == other.y
}

type head point
type tail point

func main() {
	input, err := os.ReadFile("./day09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1", "Number of unique visited points of tail", part01(string(input)))
	fmt.Println("Part 2", "Number of unique visited points of tail", part02(string(input)))
}

func part01(input string) int {
	lines := strings.Split(input, "\n")

	var moves = []move{}

	for _, line := range lines {
		if line != "" {
			moves = append(moves, getMoveFromLine(line))
		}
	}

	h := head{x: 0, y: 0}
	t := tail{x: 0, y: 0}

	var tailPositions = []point{}

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			moveHead(&h, m.direction)

			fromHead := t.distanceFromHead(h)
			if fromHead > math.Sqrt2 {
				moveTail(&t, &h)
			}
			tailPositions = append(tailPositions, point(t))
		}
	}

	var uniquePositions = []point{}

	for i, position := range tailPositions {
		unique := true
		for j := 0; j < len(uniquePositions); j++ {
			if i != j {
				if position.pointIsEqual(uniquePositions[j]) {
					unique = false
				}
			}
		}
		if unique {
			uniquePositions = append(uniquePositions, position)
		}
	}

	return len(uniquePositions)
}

func part02(input string) int {

	lines := strings.Split(input, "\n")

	var moves = []move{}

	for _, line := range lines {
		if line != "" {
			moves = append(moves, getMoveFromLine(line))
		}
	}

	h := head{
		x: 0,
		y: 0,
	}

	rope := map[int]point{}

	for i := 1; i < 10; i++ {
		rope[i] = point{
			x: 0,
			y: 0,
		}
	}

	var tailPositions = []point{}

	// adding start
	tailPositions = append(tailPositions, point{
		x: 0,
		y: 0,
	})

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			moveHead(&h, m.direction)

			for j := 1; j <= len(rope); j++ {
				t := tail(rope[j])
				var previous head
				if j == 1 {
					previous = h
				} else {
					previous = head(rope[j-1])
				}
				fromHead := t.distanceFromHead(previous)
				if fromHead > math.Sqrt2 {
					moveTail(&t, &previous)
					rope[j] = point(t)
					if j == 9 { // index 9 == tail
						tailPositions = append(tailPositions, point(t))
					}
				}
			}
		}
	}

	var uniquePositions = []point{}

	for i, position := range tailPositions {
		unique := true
		for j := 0; j < len(uniquePositions); j++ {
			if i != j {
				if position.pointIsEqual(uniquePositions[j]) {
					unique = false
				}
			}
		}
		if unique {
			uniquePositions = append(uniquePositions, position)
		}
	}

	return len(uniquePositions)
}

func moveTail(t *tail, h *head) {
	// horizontal
	if t.y == t.y && t.x != h.x {
		if t.x-h.x > 0 { // tail right of head
			t.x--
		} else {
			t.x++
		}
	}
	// vertical
	if t.x == t.x && t.y != h.y {
		if t.y-h.y > 0 { // tail above of head
			t.y--
		} else {
			t.y++
		}
	}
}

func moveHead(h *head, direction string) {
	switch direction {
	case "R":
		h.y++
	case "U":
		h.x++
	case "L":
		h.y--
	case "D":
		h.x--
	}

}

func getMoveFromLine(line string) move {
	split := strings.Split(line, " ")
	steps, _ := strconv.Atoi(split[1])
	return move{
		steps:     steps,
		direction: split[0],
	}
}
