package main

import (
	"fmt"
	_ "fmt"
	"time"
)

type State = string

func loop(state State, loopNo int) {
	fmt.Print("\033[H\033[2J")

	fmt.Println(state)
	fmt.Println("Loop #", loopNo)
	time.Sleep(1 * time.Second)
	loop(state, loopNo + 1)
}

func main() {
	loop("State1", 0)
}
