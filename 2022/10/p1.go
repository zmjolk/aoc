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

func cpuCycle(reg *int, clock chan int, total *int, instruction Instruction) {
	var alu int
	for {
		cycle := <-clock
		check(cycle, reg, total)
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

func check(cycle int, reg *int, total *int) {
	for _, v := range [6]int{20, 60, 100, 140, 180, 220} {
		if cycle == v {
			*total += (*reg * cycle)
		}
	}
}

func main() {
	in, _ := os.ReadFile("input")
	instructionSet := parse(string(in))
	clock := clock()

	reg := 1
	var total int
	for _, instruction := range instructionSet {
		cpuCycle(&reg, clock, &total, instruction)
	}
	fmt.Println(total)

}
