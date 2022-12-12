package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MonkeyOp struct {
	op  string
	arg int
}

type Monkey struct {
	items       chan int
	inspected   *int
	test        int
	trueMonkey  int
	falseMonkey int
	operation   MonkeyOp
}

func parse(in string) []Monkey {
	var monkeys []Monkey

	for _, monkSeg := range strings.Split(in, "\n\n") {
		var in int
		monkey := Monkey{items: make(chan int, 100), inspected: &in}
		lines := strings.Split(monkSeg, "\n")
		// items
		items := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		for _, item := range items {
			parsedItem, _ := strconv.Atoi(item)
			monkey.items <- parsedItem
		}

		// op
		op := strings.Split(strings.Split(lines[2], "old ")[1], " ")
		monkey.operation.op = op[0]
		arg, err := strconv.Atoi(op[1])
		if err != nil {
			monkey.operation.op = "**"
		} else {
			monkey.operation.arg = arg
		}

		// test
		test, _ := strconv.Atoi(strings.Split(lines[3], "by ")[1])
		monkey.test = test

		// true & false
		trueMonkey, _ := strconv.Atoi(strings.Split(lines[4], "monkey ")[1])
		monkey.trueMonkey = trueMonkey

		falseMonkey, _ := strconv.Atoi(strings.Split(lines[5], "monkey ")[1])
		monkey.falseMonkey = falseMonkey

		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func MonkeyTime(monkeys []Monkey) {
	for _, monkey := range monkeys {
		for {
			// var item int
			if len(monkey.items) == 0 {
				break
			}
			item := <-monkey.items
			*monkey.inspected++
			switch monkey.operation.op {
			case "*":
				item *= monkey.operation.arg
			case "+":
				item += monkey.operation.arg
			case "**":
				item = item * item
			}
			item /= 3

			var targetMonkey int
			if item%monkey.test == 0 {
				targetMonkey = monkey.trueMonkey
			} else {
				targetMonkey = monkey.falseMonkey
			}
			monkeys[targetMonkey].items <- item
		}
	}
}

func totalItems(monkeys []Monkey) int {
	var total int
	for _, monkey := range monkeys {
		total += len(monkey.items)
	}
	return total
}

func main() {
	in, _ := os.ReadFile("input")
	monkeys := parse(string(in))

	for i := 0; i < 20; i++ {
		MonkeyTime(monkeys)
	}

	var inspected []int
	for _, monkey := range monkeys {
		inspected = append(inspected, *monkey.inspected)
	}
	sort.Ints(inspected)

	// MonkeyBusiness
	fmt.Println(inspected[len(inspected)-1] * inspected[len(inspected)-2])

	// fmt.Printf("%#v", monkeys)
}
