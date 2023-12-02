package main

import (
	"github.com/atomragnar/advent-of-code-2023/day/one"
	"github.com/atomragnar/advent-of-code-2023/day/two"
)

type adventFunc func(bool)

var adventFuncs = map[int]adventFunc{
	1: one.Solution,
	2: two.Solution,
}

func main() {
	adventFuncs[2](true)
}
