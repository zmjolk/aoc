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
	sacks := strings.Split(string(in), "\n")
	var result int

	for i := 3; i < len(sacks); i+=3 {

		sackSquad := sacks[i-3:i]

		var masks [3]int
		for idx, sack := range(sackSquad) {
			for _, char := range(sack) {
				masks[idx] = masks[idx] | (0b1 << mF[char])
			}
		}
		maskResult := masks[0] & masks[1] & masks[2]
		item := bits.Len64(uint64(maskResult)) +1
		result += item
	}
	fmt.Println(result)
}
