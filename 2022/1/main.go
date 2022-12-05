package main

import(
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	in, _ := os.ReadFile("input")
	var top []int
	var reg int
	for _, s := range(strings.Split(string(in),"\n")) {
		res, _ := strconv.Atoi(s)
		reg += res
		if s == "" {
			top = append(top, reg)
			reg = 0
		}
	}
	sort.Sort(sort.IntSlice(top))
	
	for _, x := range(top[len(top)-3:]) {
		reg += x
	}
	fmt.Println(reg)
}
