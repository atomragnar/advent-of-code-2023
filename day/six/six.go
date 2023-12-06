package six

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"log"
	"math"
	"strconv"
	"strings"
)

func Solution(isPartTwo bool) {
	dataString := `Time:        38     94     79     70
Distance:   241   1549   1074   1091`
	var fn func(string)
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	fn(dataString)
}

func partTwo(input string) {
	var result float64
	// (t-x)x > r
	d := strings.Split(input, "\n")
	td := strings.TrimSpace(strings.Split(d[0], ":")[1])
	rd := strings.TrimSpace(strings.Split(d[1], ":")[1])
	var ts string
	var rs string

	for i := 0; i < len(td); i++ {
		if util.IsDigit(td[i]) {
			ts += string(td[i])
		}
	}
	for i := 0; i < len(rd); i++ {
		if util.IsDigit(rd[i]) {
			rs += string(rd[i])
		}
	}

	tf, err := strconv.ParseFloat(ts, 64)
	if err != nil {
		log.Fatal(err)
	}

	rf, err := strconv.ParseFloat(rs, 64)
	if err != nil {
		log.Fatal(err)
	}

	sqt := math.Pow(tf, 2)
	r4 := 4 * rf
	b := math.Sqrt(sqt - r4)
	bl := (tf + b) / 2
	bh := math.Floor((tf - b) / 2)
	res := math.Ceil(bl-bh) - 1
	result = res

	fmt.Printf("Results is: %d\n", int(result))
}

func partOne(input string) {
	var result int
	result = 1

	d := strings.Split(input, "\n")
	td := strings.TrimSpace(strings.Split(d[0], ":")[1])
	rd := strings.TrimSpace(strings.Split(d[1], ":")[1])
	t := util.StringToInts(td)
	r := util.StringToInts(rd)

	var c int
	var wins int

	fmt.Println(t)
	fmt.Println(r)

	for i, time := range t {
		wins = 0
		c = 0
		for c <= r[i] {
			if (time-c)*c > r[i] {
				wins++
			}
			c++
		}
		result *= wins
	}

	fmt.Printf("Results is: %d", result)
}
