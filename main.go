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

func NewState(size int) *State {
	cells := make([]bool, size*size)
	return &State{cells: cells, size: size}
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
	s.cells[s.Index(x, y)] = state
}

func (s State) Get(x int, y int) bool {
	return s.cells[s.Index(x, y)]
}

func (s State) Print() {
	for x := 0; x < s.size; x++ {
		for y := 0; y < s.size; y++ {
			if s.Get(x, y) {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func loop(state *State, loopNo int) {
	fmt.Print("\033[H\033[2J")

	state.Print()
	fmt.Println("Loop #", loopNo)

	time.Sleep(1 * time.Second)
	loop(state, loopNo+1)
}

func main() {
	const size = 16
	s := NewState(size)
	loop(s, 0)
}
