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

func (f Forest) checkTreeViewDist(t Tree, direction string) int {

	for i := 1; true; i++ {
		if direction == "n" {
			if i > t.y {
				return i - 1
			} else if t.h <= f.trees[t.y-i][t.x].h {
				return i
			}
		}
		if direction == "e" {
			if i >= len(f.trees[0])-t.x {
				return i - 1
			} else if t.h <= f.trees[t.y][t.x+i].h {
				return i
			}
		}
		if direction == "s" {
			if i >= len(f.trees)-t.y {
				return i - 1
			} else if t.h <= f.trees[t.y+i][t.x].h {
				return i
			}
		}
		if direction == "w" {
			if i > t.x {
				return i - 1
			} else if t.h <= f.trees[t.y][t.x-i].h {
				return i
			}
		}
	}
	return 0
}

func calcViewScore(viewDists []int) int {
	viewScore := 1
	for _, v := range viewDists {
		viewScore *= v
	}
	return viewScore
}

func main() {
	in, _ := os.ReadFile("input")
	cardinals := [4]string{"n", "e", "s", "w"}
	f := buildForest(string(in))

	var highest int
	for _, treeList := range f.trees {
		for _, t := range treeList {
			var viewDists []int
			for _, dir := range cardinals {
				viewDists = append(viewDists, f.checkTreeViewDist(t, dir))
			}
			viewScore := calcViewScore(viewDists)
		}
	}
	fmt.Println("highest is", highest)
}
