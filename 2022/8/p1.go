package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Forest struct {
	trees [][]Tree
}

type Tree struct {
	x, y, h int
}

func buildForest(in string) Forest {
	f := Forest{}
	for y, line := range strings.Split(in, "\n") {
		var row []Tree
		for x, char := range []rune(line) {
			h, _ := strconv.Atoi(string(char))
			row = append(row, Tree{h: h, x: x, y: y})
		}
		f.trees = append(f.trees, row)
	}
	return f
}

func (f Forest) checkTreeVisible(t Tree, direction string) bool {

	for i := 1; true; i++ {
		if direction == "n" {
			if i > t.y {
				return true
			}
			if t.h <= f.trees[t.y-i][t.x].h {
				return false
			}
		}
		if direction == "e" {
			if i >= len(f.trees[0])-t.x {
				return true
			}
			if t.h <= f.trees[t.y][t.x+i].h {
				return false
			}
		}
		if direction == "s" {
			if i >= len(f.trees)-t.y {
				return true
			}
			if t.h <= f.trees[t.y+i][t.x].h {
				return false
			}
		}
		if direction == "w" {
			if i > t.x {
				return true
			}
			if t.h <= f.trees[t.y][t.x-i].h {
				return false
			}
		}
	}
	return false
}

func main() {
	in, _ := os.ReadFile("input")
	cardinals := [4]string{"n", "e", "s", "w"}
	f := buildForest(string(in))

	var visible int
	for _, treeList := range f.trees {
		for _, t := range treeList {
			var tVis bool
			for _, dir := range cardinals {
				if f.checkTreeVisible(t, dir) {
					tVis = true
				}
			}
			if tVis {
				visible++
			}
		}
	}
	fmt.Println(visible)
}
