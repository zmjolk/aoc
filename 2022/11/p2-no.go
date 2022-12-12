// This technically works but is quite computationally spicy

package main

import (
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MonkeyOp struct {
	op  string
	arg *big.Int
}

type Monkey struct {
	items       chan *big.Int
	inspected   *int
	test        *big.Int
	trueMonkey  int
	falseMonkey int
	operation   MonkeyOp
}

func parse(in string) []Monkey {
	var monkeys []Monkey

	for _, monkSeg := range strings.Split(in, "\n\n") {
		var in int
		monkey := Monkey{items: make(chan *big.Int, 500), inspected: &in}
		lines := strings.Split(monkSeg, "\n")
		// items
		items := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		for _, item := range items {
			parsedItem, _ := strconv.Atoi(item)
			monkey.items <- big.NewInt(int64(parsedItem))
		}

		// op
		op := strings.Split(strings.Split(lines[2], "old ")[1], " ")
		monkey.operation.op = op[0]
		arg, err := strconv.Atoi(op[1])
		if err != nil {
			monkey.operation.op = "**"
		} else {
			monkey.operation.arg = big.NewInt(int64(arg))
		}

		// test
		test, _ := strconv.Atoi(strings.Split(lines[3], "by ")[1])
		monkey.test = big.NewInt(int64(test))

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
			// fmt.Println("Doing maths on ", item.String())
			*monkey.inspected++
			switch monkey.operation.op {
			case "*":
				item.Mul(item, monkey.operation.arg)
			case "+":
				item.Add(item, monkey.operation.arg)
			case "**":
				item.Mul(item, item)
			}
			// item /= 3

			var targetMonkey int
			mod := new(big.Int)
			mod.Mod(item, monkey.test)
			if mod.Int64() == 0 {
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

	for i := 0; i < 10000; i++ {
		fmt.Println("MonkeyBusiness round: ", i)
		MonkeyTime(monkeys)
	}

	var inspected []int
	for _, monkey := range monkeys {
		inspected = append(inspected, *monkey.inspected)
	}
	sort.Ints(inspected)

	// MonkeyBusiness
	fmt.Println(inspected[len(inspected)-1] * inspected[len(inspected)-2])
}
