package main

import (
	"fmt"
	"os"
	// "regexp"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	label rune
}

type Tower struct {
	boxes []Box
}

type CargoYard struct {
	towers []Tower
}

type Instruction struct {
	move, from, to int
}

func reverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func parse(in string) (*CargoYard, []Instruction) {
	split := strings.Split(in, "\n\n")
	initConfig := strings.Split(split[0], "\n")
	reverseSlice(initConfig)

	var cargoYard CargoYard
	for i := 0; i < (len(initConfig[0])+1)/4; i++ {
		cargoYard.towers = append(cargoYard.towers, Tower{})
	}
	initConfig = initConfig[1:]

	for _, v := range initConfig {
		for j := 0; j < (len(initConfig[0])+1)/4; j++ {
			label := []rune(v)[j+1+j*3]
			if label == []rune(" ")[0] {
				continue
			}
			cargoYard.towers[j].boxes = append(cargoYard.towers[j].boxes, Box{label: label})
		}
	}
	var instructions []Instruction

	for _, v := range strings.Split(split[1], "\n") {
		instruction := strings.Split(v, " ")
		move, _ := strconv.Atoi(instruction[1])
		from, _ := strconv.Atoi(instruction[3])
		to, _ := strconv.Atoi(instruction[5])
		instructions = append(instructions, Instruction{move, from, to})
	}
	return &cargoYard, instructions
}

func (cargoYard *CargoYard) activateCrane(instructions []Instruction) {
	for _, instruction := range instructions {

		from := &cargoYard.towers[instruction.from-1]
		to := &cargoYard.towers[instruction.to-1]
		to.boxes = append(to.boxes, from.boxes[len(from.boxes)-instruction.move:]...)
		from.boxes = from.boxes[:len(from.boxes)-instruction.move]
	}
}

func (cargoYard *CargoYard) checkTopBoxes() string {
	var result []rune
	for _, tower := range cargoYard.towers {
		result = append(result, tower.boxes[len(tower.boxes)-1].label)
	}
	return string(result)
}

func main() {
	in, _ := os.ReadFile("input")
	cargoYard, instructions := parse(string(in))

	cargoYard.activateCrane(instructions)
	fmt.Println(cargoYard.checkTopBoxes())
}
