package main

import (
	"fmt"
	"os"
)

type RadioReceiver struct {
	quit  chan rune
	data  chan rune
	count int
}

func makeRadioReceiver(in string) *RadioReceiver {
	r := RadioReceiver{
		quit: make(chan rune, 1),
		data: make(chan rune),
	}
	go func(x []rune, r *RadioReceiver) {
		for _, y := range x {
			select {
			case <-r.quit:
				return
			default:
				r.count++
				r.data <- y
			}
		}
	}([]rune(in), &r)
	return &r
}

func radioBufferUnique(b []rune) bool {
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] == b[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	in, _ := os.ReadFile("input")
	radioReceiver := makeRadioReceiver(string(in))
	fmt.Printf("rec %T", radioReceiver.count)

	buf := []rune{
		<-radioReceiver.data,
		<-radioReceiver.data,
		<-radioReceiver.data,
	}
	for {
		buf = append(buf, <-radioReceiver.data)
		if radioBufferUnique(buf) {
			fmt.Println("count: ", radioReceiver.count)
			radioReceiver.quit <- <-radioReceiver.data
			return
		}
		buf = buf[1:]
	}
}
