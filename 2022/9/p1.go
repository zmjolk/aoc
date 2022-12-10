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
	head, tail RopeComponent
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
		// move head
		switch instruction.direction {
		case North:
			r.head.y += 1
		case East:
			r.head.x += 1
		case South:
			r.head.y -= 1
		case West:
			r.head.x -= 1
		}
		// check tail movement
		xDist := r.head.x - r.tail.x
		yDist := r.head.y - r.tail.y

		if math.Abs(float64(xDist)) > 1 {
			fmt.Println("X Dist needs moving!", r.head, r.tail)
			if r.head.x > r.tail.x {
				r.tail.x = r.tail.x + xDist - 1
			} else {
				r.tail.x = r.tail.x + xDist + 1
			}
			if math.Abs(float64(yDist)) > 0 {
				r.tail.y = r.head.y
			}
			fmt.Println("X Dist moved!", r.head, r.tail)
		}
		if math.Abs(float64(yDist)) > 1 {
			fmt.Println("Y Dist needs moving!", r.head, r.tail)
			if r.head.y > r.tail.y {
				r.tail.y = r.tail.y + yDist - 1
			} else {
				r.tail.y = r.tail.y + yDist + 1
			}
			if math.Abs(float64(xDist)) > 0 {
				r.tail.x = r.head.x
			}
			fmt.Println("Y Dist moved!", r.head, r.tail)
		}

		// update visited map
		if _, ok := g.square[r.tail.x]; !ok {
			g.square[r.tail.x] = make(map[int]Visited)
		}
		if _, ok := g.square[r.tail.x][r.tail.y]; !ok {
			g.square[r.tail.x][r.tail.y] = true
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

func main() {
	in, _ := os.ReadFile("input")
	instructionSet := parse(string(in))

	grid := Grid{
		square: map[int]map[int]Visited{
			0: map[int]Visited{
				0: true,
			},
		},
	}
	rope := Rope{
		head: RopeComponent{x: 0, y: 0},
		tail: RopeComponent{x: 0, y: 0},
	}

	for i, v := range instructionSet {
		fmt.Printf("executing: %#v\n", v)
		rope.executeInstruction(v, grid)
		if i > 20 {
			// break
		}
	}
	// fmt.Println(grid)
	fmt.Println(calcVisited(grid))
}
