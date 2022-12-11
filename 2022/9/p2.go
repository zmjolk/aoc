package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int
type Visited bool

const (
	North Direction = iota
	East            = iota
	South           = iota
	West            = iota
)

type Grid struct {
	square map[int]map[int]Visited
}

type Rope struct {
	knots []RopeComponent
}

type RopeComponent struct {
	x, y int
}

type Instruction struct {
	iterations int
	direction  Direction
}

func parse(in string) []Instruction {
	var instructionSet []Instruction
	for _, v := range strings.Split(in, "\n") {
		set := strings.Split(v, " ")
		iterations, _ := strconv.Atoi(set[1])
		var direction Direction
		switch set[0] {
		case "U":
			direction = North
		case "R":
			direction = East
		case "L":
			direction = West
		case "D":
			direction = South
		}
		instructionSet = append(instructionSet, Instruction{iterations: iterations, direction: direction})
	}
	return instructionSet
}

func (r *Rope) executeInstruction(instruction Instruction, g Grid) {
	for i := 0; i < instruction.iterations; i++ {
		for j := 0; j < len(r.knots)-1; j++ {
			knot := &r.knots[j]
			nextKnot := &r.knots[j+1]

			// move head if its head time
			if j == 0 {
				switch instruction.direction {
				case North:
					knot.y += 1
				case East:
					knot.x += 1
				case South:
					knot.y -= 1
				case West:
					knot.x -= 1
				}
			}
			// check tail movement
			xDist := knot.x - nextKnot.x
			yDist := knot.y - nextKnot.y

			if math.Abs(float64(xDist)) > 1 {
				if knot.x > nextKnot.x {
					nextKnot.x = nextKnot.x + 1
				} else {
					nextKnot.x = nextKnot.x - 1
				}
				if math.Abs(float64(yDist)) > 0 {
					if knot.y > nextKnot.y {
						nextKnot.y = nextKnot.y + 1
					} else {
						nextKnot.y = nextKnot.y - 1
					}

				}
			} else if math.Abs(float64(yDist)) > 1 {
				if knot.y > nextKnot.y {
					nextKnot.y = nextKnot.y + 1
				} else {
					nextKnot.y = nextKnot.y - 1
				}
				if math.Abs(float64(xDist)) > 0 {
					if knot.x > nextKnot.x {
						nextKnot.x = nextKnot.x + 1
					} else {
						nextKnot.x = nextKnot.x - 1
					}
				}
			}
		}
		finalKnot := r.knots[len(r.knots)-1]
		// update visited map for final knot
		if _, ok := g.square[finalKnot.x]; !ok {
			g.square[finalKnot.x] = make(map[int]Visited)
		}
		if _, ok := g.square[finalKnot.x][finalKnot.y]; !ok {
			g.square[finalKnot.x][finalKnot.y] = true
		}
	}
}

func calcVisited(g Grid) int {
	var total int
	for _, d := range g.square {
		total += len(d)
	}
	return total
}

func Visualise(g Grid) {
	leng := 75
	for i := -leng; i < leng; i++ {
		for j := -leng; j < leng; j++ {
			char := "."
			if g.square[i][j] {
				char = "#"
			}
			fmt.Printf("%v", char)
		}
		fmt.Printf("\n")
	}
}

func main() {
	in, _ := os.ReadFile("input")
	// in, _ := os.ReadFile("itest")
	instructionSet := parse(string(in))

	grid := Grid{
		square: map[int]map[int]Visited{
			0: map[int]Visited{
				0: true,
			},
		},
	}
	rope := Rope{}
	for i := 0; i < 10; i++ {
		rope.knots = append(rope.knots, RopeComponent{})
	}

	for i, v := range instructionSet {
		// fmt.Printf("executing: %#v\n", v)
		rope.executeInstruction(v, grid)
		if i > 18 {
			// break
		}
	}
	Visualise(grid)
	fmt.Println(calcVisited(grid))
}
