package main

import (
	"github.com/atomragnar/advent-of-code-2023/day/four"
	"github.com/atomragnar/advent-of-code-2023/day/one"
	"github.com/atomragnar/advent-of-code-2023/day/three"
	"github.com/atomragnar/advent-of-code-2023/day/two"
)

type adventFunc func(bool)

var adventFuncs = map[int]adventFunc{
	1: one.Solution,
	2: two.Solution,
	3: three.Solution,
	4: four.Solution,
}

func main() {
	adventFuncs[4](false)
}
