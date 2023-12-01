package main

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/one"
)

type adventFunc func(bool)

var adventFuncs = map[int]adventFunc{
	1: one.Solution,
}

func main() {
	adventFuncs[1](false)
	fmt.Print("\n")
	adventFuncs[1](true)
}
