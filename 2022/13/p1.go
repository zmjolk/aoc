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
	children   []*Packet
}

const (
	ARRSTART     = '['
	ARREND       = ']'
	ARRSEP       = ','
	TRUE     int = iota
	FALSE
	CONTINUE
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

func parse(in string) (parsed [][2]*Packet) {
	for _, packetStringTups := range strings.Split(in, "\n\n") {
		packetStrings := strings.Split(packetStringTups, "\n")
		var packetTup [2]*Packet
		packetTup[0] = parsePacket(packetStrings[0])
		packetTup[1] = parsePacket(packetStrings[1])
		parsed = append(parsed, packetTup)
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

	var total int
	for i, packet := range packets {
		if ok := process(packet); ok == TRUE {
			fmt.Println("packet ", i+1, "TRUE")
			total += i + 1
		} else if ok == CONTINUE {
			fmt.Println("???")
		} else if ok == FALSE {
			fmt.Println("packet", i+1, "FALSE")
		}
	}
	fmt.Println(total)
}
