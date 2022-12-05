package main

import(
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
)

func aContainsB(x ...int) bool {
	fmt.Println(x)
	if x[0] < x[2] && x[0] < x[2] ||
	   x[0] > x[3] && x[0] > x[3] {
		return false
	}  
	return true
}

func main() {


	in, _ := os.ReadFile("input")
	re := regexp.MustCompile(`\d+`)
	var result int

	for _, pair := range(strings.Split(string(in), "\n")) {
		var aF []int
		for _, x := range(re.FindAllString(pair, 4)) {
			i, _ := strconv.Atoi(x)
			aF = append(aF, i)
		}
		aB := []int{aF[2], aF[3], aF[0], aF[1]}
		
		if aContainsB(aF...) || aContainsB(aB...) {
			result ++
		}
	}
	fmt.Println(result)
	
}
