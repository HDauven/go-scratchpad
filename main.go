package main

import (
	"fmt"

	"github.com/HDauven/go-scratchpad/hof"
)

func main() {
	useTimes()
}

func useTimes() {
	double := hof.Times(2)
	fmt.Println(double(5))
}
