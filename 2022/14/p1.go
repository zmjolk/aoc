package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cave struct {
	points [][]Point
	sands  int
}

const (
	FREE int = iota
	ROCK
	SAND
)

type Point struct {
	x, y, obj int
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func cmp(a int, b int) int {
	if a > b {
		return -1
	}
	return 1
}

func (c *Cave) fillRock(a *Point, b *Point) {
	if a.x != b.x {
		for i := 0; i <= int(math.Abs(float64(a.x)-float64(b.x))); i++ {
			c.points[a.x+i*cmp(a.x, b.x)][a.y].obj = ROCK
		}
	} else if a.y != b.y {
		for i := 0; i <= int(math.Abs(float64(a.y)-float64(b.y))); i++ {
			c.points[a.x][a.y+i*cmp(a.y, b.y)].obj = ROCK
		}
	}
}

func parse(in string) Cave {
	c := Cave{points: make([][]Point, 250)}
	for i := 0; i < 250; i++ {
		c.points[i] = make([]Point, 250)
		for j := 0; j < 250; j++ {
			c.points[i][j] = Point{x: i, y: j, obj: FREE}
		}
	}
	for _, rockInstructions := range strings.Split(in, "\n") {
		var curPoint, prevPoint *Point
		for _, rockInstruction := range strings.Split(rockInstructions, " -> ") {
			re := regexp.MustCompile(`(\d+),(\d+)`)
			matches := re.FindStringSubmatch(rockInstruction)
			curX, curY := parseInt(matches[1]), parseInt(matches[2])
			curX = curX - 450

			prevPoint = curPoint
			curPoint = &c.points[curX][curY]

			if prevPoint != nil {
				c.fillRock(prevPoint, curPoint)
			}
		}
	}
	return c
}

func (c *Cave) visualize() {

	for i := 0; i < 165; i++ {
		for j := 0; j < 100; j++ {
			if c.points[j][i].obj == FREE {
				fmt.Printf(".")
			} else if c.points[j][i].obj == ROCK {
				fmt.Printf("#")
			} else if c.points[j][i].obj == SAND {
				fmt.Printf("o")
			}
		}
		fmt.Printf("\n")
	}
}

func (c *Cave) dropSand() bool {
	c.points[50][0].obj = SAND

	var curPoint, nextPoint *Point
	curPoint = &c.points[50][0]

	for {
		if curPoint.y == len(c.points[0])-1 {
			return false
		} else if next := c.points[curPoint.x][curPoint.y+1]; next.obj == FREE {
			nextPoint = &next
		} else if next := c.points[curPoint.x-1][curPoint.y+1]; next.obj == FREE {
			nextPoint = &next
		} else if next := c.points[curPoint.x+1][curPoint.y+1]; next.obj == FREE {
			nextPoint = &next
		} else {
			return true
		}
		c.points[curPoint.x][curPoint.y].obj = FREE
		c.points[nextPoint.x][nextPoint.y].obj = SAND
		curPoint = nextPoint
	}
}

func main() {
	in, _ := os.ReadFile("input")
	c := parse(string(in))
	_ = c

	for {
		if ok := c.dropSand(); !ok {
			break
		} else {
			c.sands++
		}
	}
	c.visualize()
	fmt.Println(c.sands)
}
