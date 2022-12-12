package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NOOP = "noop"
	ADD  = "addx"
)

type Instruction struct {
	op  string
	arg int
}

func parse(in string) []Instruction {
	var instructionSet []Instruction
	for _, line := range strings.Split(in, "\n") {
		lineSegs := strings.Split(line, " ")
		if lineSegs[0] == NOOP {
			instructionSet = append(instructionSet, Instruction{op: NOOP})
		} else {
			arg, _ := strconv.Atoi(lineSegs[1])
			instructionSet = append(instructionSet, Instruction{op: ADD, arg: arg})
		}
	}
	return instructionSet
}

func clock() chan int {
	ticker := make(chan int)
	go func(t chan int) {
		for i := 1; true; i++ {
			t <- i
		}
	}(ticker)
	return ticker
}

func cpuCycle(reg *int, clock chan int, instruction Instruction) {
	var alu int
	for {
		cycle := <-clock
		draw(cycle, reg)
		if instruction.op == NOOP {
			return
		} else {
			if alu != 0 {
				*reg += alu
				return
			} else {
				alu = instruction.arg
			}
		}
	}
}

func draw(cycle int, reg *int) {
	position := cycle - 1
	if position >= 40 {
		position = (cycle - 1) % 40
	}
	pixel := "."
	if position == *reg-1 || position == *reg || position == *reg+1 {
		pixel = "#"
	}
	_ = pixel
	if position == 0 {
		fmt.Printf("\n")
	}
	fmt.Printf("%v", pixel)
}

func main() {
	in, _ := os.ReadFile("input")
	instructionSet := parse(string(in))
	clock := clock()

	reg := 1
	for _, instruction := range instructionSet {
		cpuCycle(&reg, clock, instruction)
	}
}
