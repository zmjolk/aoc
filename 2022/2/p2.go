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

var q = [5]int{
	3,
	1,
	2,
	3,
	1,
}

func game(p int, o int) (myPlay int) {
	if p == 1 {
		o--
	} else if p == 3 {
		o++
	}
	return q[o]
}

func main() {
	var score int
	in, _ := os.ReadFile("input")
	for _, round := range bytes.Split(in, []byte("\n")) {
		plays := strings.Split(string(round), " ")
		score += game(m[plays[1]], m[plays[0]]) + (m[plays[1]]-1)*3
	}
	fmt.Println(score)
}
