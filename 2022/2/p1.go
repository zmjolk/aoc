package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	lose = iota
	draw = iota * 3
	win  = iota * 3
)

var m = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func game(p int, o int) (points int) {
	if p-o == 1 || p-o == -2 {
		return win
	} else if o == p {
		return draw
	}
	return lose
}

func main() {
	var score int
	in, _ := os.ReadFile("input")
	for _, round := range bytes.Split(in, []byte("\n")) {
		plays := strings.Split(string(round), " ")
		score += m[plays[1]] + game(m[plays[1]], m[plays[0]])
	}
	fmt.Println(score)
}
