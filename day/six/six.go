package six

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solution(isPartTwo bool) {
	day := 6
	data := func() ([]byte, error) {
		d, err := os.ReadFile(util.DataPath(day))
		if err != nil {
			return nil, err
		}
		return d, err
	}
	var fn func(func() ([]byte, error))
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	fn(data)
}

func partTwo(input func() ([]byte, error)) {
	var result float64

	data, _ := input()

	d := util.SplitData(data)
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

func partOne(input func() ([]byte, error)) {
	var result int
	result = 1

	data, _ := input()

	d := util.SplitData(data)
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
