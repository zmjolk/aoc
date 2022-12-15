package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	value      int
	packetType string
	divider    int
	children   []*Packet
}

const (
	TRUE int = iota
	FALSE
	CONTINUE
	ARRSTART = '['
	ARREND   = ']'
	ARRSEP   = ','
)

func scanInt(x string) int {
	i, _ := strconv.Atoi(x)
	return i
}

func parsePacket(in string) *Packet {
	p := &Packet{packetType: "list", children: []*Packet{}}
	for i := 1; i < len(in); i++ {
		if in[i] == ARRSTART {
			open := 1
			for j := i + 1; j < len(in); j++ {
				if in[j] == ARRSTART {
					open++
				}
				if in[j] == ARREND {
					open--
				}
				if open == 0 {
					p.children = append(p.children, parsePacket(in[i:j+1]))
					i = j
					break
				}
			}
		} else if in[i] == ARRSEP {
			continue
		} else {
			for j := i + 1; j < len(in); j++ {
				if in[j] == ARRSEP || in[j] == ARREND {
					pChild := Packet{packetType: "value", value: scanInt(in[i:j])}
					p.children = append(p.children, &pChild)
					i = j
					break
				}
			}
		}
	}
	return p
}

func parse(in string) (parsed []*Packet) {
	for _, packetStringTups := range strings.Split(in, "\n\n") {
		for _, v := range strings.Split(packetStringTups, "\n") {
			parsed = append(parsed, parsePacket(string(v)))
		}
	}
	return parsed
}

func compare(a, b int) int {
	if a < b {
		return TRUE
	}
	if a > b {
		return FALSE
	}
	return CONTINUE
}

func process(p [2]*Packet) int {
	left := p[0]
	right := p[1]
	if left.packetType == "value" && right.packetType == "value" {
		return compare(left.value, right.value)
	}
	if left.packetType == "value" {
		left.packetType = "list"
		left.children = []*Packet{{packetType: "value", value: left.value}}
	}
	if right.packetType == "value" {
		right.packetType = "list"
		right.children = []*Packet{{packetType: "value", value: right.value}}
	}
	smaller := len(left.children)
	if len(right.children) < smaller {
		smaller = len(right.children)
	}

	for i := 0; i < smaller; i++ {
		check := process([2]*Packet{left.children[i], right.children[i]})
		if check == CONTINUE {
			continue
		}
		return check
	}
	return compare(len(left.children), len(right.children))
}

func main() {
	in, _ := os.ReadFile("input")
	packets := parse(string(in))
	fmt.Println(packets)

	dividerTwo := Packet{
		packetType: "list",
		divider:    2,
		children: []*Packet{{
			packetType: "list",
			children: []*Packet{{
				packetType: "value",
				value:      2,
			}},
		}},
	}
	dividerSix := Packet{
		packetType: "list",
		divider:    6,
		children: []*Packet{{
			packetType: "list",
			children: []*Packet{{
				packetType: "value",
				value:      6,
			}},
		}},
	}
	packets = append(packets, &dividerTwo)
	packets = append(packets, &dividerSix)

	for i := 0; i < len(packets); i++ {
		for j := i + 1; j < len(packets); j++ {
			proc := process([2]*Packet{packets[i], packets[j]})
			if proc != TRUE {
				packets[i], packets[j] = packets[j], packets[i]
			}

		}
	}

	var twoIdx, sixIdx int
	for i, p := range packets {
		if p.divider == 6 {
			sixIdx = i + 1
		}
		if p.divider == 2 {
			twoIdx = i + 1
		}
	}
	fmt.Println(twoIdx * sixIdx)
}
