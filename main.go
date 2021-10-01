package main

import (
	"fmt"
	_ "fmt"
	"time"
)

type State struct {
	cells []bool // Flattened size x size array
	size  int
}

func (s State) Index(x int, y int) int {
	if x < 0 {
		x = s.size - 1
	} else if x >= s.size {
		x = 0
	}
	if y < 0 {
		y = s.size - 1
	} else if y >= s.size {
		y = 0
	}
	return x*s.size + y
}

func (s State) Set(x int, y int, state bool) {
	const z = 0

}

func (s State) Get(x int, y int) bool {
	return true
}

// Rules
// - Is dead & = 3 live neighbors --> Live
// - Is live & < 2 live neighbors --> Dies
// - Is live & 2 | 3 live neighbors --> Live
// - Is live & > 3 live neighbors --> Dies

// X X X
// X O X
// X X X

func loop(state State, loopNo int) {
	fmt.Print("\033[H\033[2J")

	fmt.Println(state)
	fmt.Println("Loop #", loopNo)

	time.Sleep(1 * time.Second)
	loop(state, loopNo+1)
}

func main() {
	loop(_, 0)
}
