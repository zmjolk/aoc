package main

import(
	"fmt"
	"os"
	"math/bits"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	mF := make(map[rune]int)
	mB := make(map[int]rune)
	for i, c := range(charset) {
		mF[c+1]= i
		mB[i+1]= c
	}

	in, _ := os.ReadFile("input")
	var result int

	for _, sack := range(strings.Split(string(in), "\n")) {
		first := sack[:len(sack)/2]
		second := sack[len(sack)/2:]

		var maskFirst, maskSecond int
		var item int
		for idx, _ := range(first) {
			maskFirst = maskFirst | (0b1 << mF[rune(first[idx])])
			maskSecond = maskSecond | (0b1 << mF[rune(second[idx])])
		}
		item = maskFirst & maskSecond
		char := bits.Len64(uint64(item))
		result += char
	}
	fmt.Println(result)
}
