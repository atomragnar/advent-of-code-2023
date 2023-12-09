package three

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"io"
	"log/slog"
	"strconv"
)

// this solution became a mess, but it works

func Solution(isPartTwo bool) {
	day := 3
	var fn util.BufferProcessor
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	err := util.ProcessInput(util.DataPath(day), fn)
	if err != nil {
		return
	}
}

type Line struct {
	numbers map[int]int64
	symbols map[int]struct{}
}

func newLine() *Line {
	return &Line{
		numbers: make(map[int]int64),
		symbols: make(map[int]struct{}),
	}
}

func numFromString(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		slog.Error("Error parsing number", "error", err)
		return 0
	}
	return num
}

func (l *Line) addNum(n string, col int) {
	num := numFromString(n)
	for i := 0; i < len(n); i++ {
		l.numbers[col+i] = num
	}
}

func (l *Line) addSymbol(col int) {
	l.symbols[col] = struct{}{}
}

func (l *Line) hasNumber(col int) bool {
	_, ok := l.numbers[col]
	return ok
}

func (l *Line) getNum(col int) int64 {
	return l.numbers[col]
}

func (l *Line) empty() bool {
	return len(l.numbers) == 0 && len(l.symbols) == 0
}

func (l *Line) forSymbol(fn func(col int)) {
	for col := range l.symbols {
		fn(col)
	}
}

type Lines struct {
	current *Line
	prev    *Line
}

func newLines() *Lines {
	return &Lines{
		current: newLine(),
		prev:    newLine(),
	}
}

func (l *Lines) update() {
	l.prev = l.current
	l.current = newLine()
}

func (l *Lines) prevEmpty() bool {
	return l.prev.empty()
}

func (l *Lines) compare() int64 {
	return sumNumbers(l.prev, l.current) + sumNumbers(l.current, l.prev)
}

func sumNumbers(l1 *Line, l2 *Line) int64 {
	var sum int64
	l1.forSymbol(func(col int) {
		numbers := make(map[int64]struct{})
		for i := -1; i < 2; i++ {
			if l2.hasNumber(col + i) {
				numbers[l2.getNum(col+i)] = struct{}{}
			}
		}
		for num := range numbers {
			sum += num
		}
	})
	return sum
}

func partTwo(reader *bufio.Reader) error {
	var result int64
	result = 0
	var l string
	var start int
	var end int
	var star string
	var index int
	lines := newLines()
	starMap := make(map[string][]int64)
	stars := make(map[int][]string)
	nLine := 1
	for {

		chunk, err := reader.ReadBytes('\r')
		stars[nLine] = make([]string, 0)

		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		l = string(chunk)
		for i := 0; i < len(l); i++ {

			if l[i] != '*' && !util.IsDigit(l[i]) {
				continue
			}

			star = fmt.Sprintf("%d,%d", i, nLine)
			stars[nLine] = append(stars[nLine], star)
			starMap[star] = make([]int64, 0)

			if util.IsDigit(l[i]) {
				start, end = i, i

				for {
					i++
					if i == len(l) {
						break
					}
					if !util.IsDigit(l[i]) {
						end = i
						i--
						break
					}
				}

				lines.current.addNum(l[start:end], start)
				continue
			}

			lines.current.addSymbol(i)

			if i-1 >= 0 && util.IsDigit(l[i-1]) {
				index = i - 1
				end = index
				for {
					if index == 0 {
						break
					}
					if !util.IsDigit(l[index]) {
						start = index
						break
					}
					index--
				}
				starMap[star] = append(starMap[star], numFromString(l[start+1:end+1]))
			}

			if i+1 < len(l) && util.IsDigit(l[i+1]) {
				index = i + 1
				start = index
				for {
					if index == len(l) {
						break
					}
					if !util.IsDigit(l[index]) {
						end = index
						break
					}
					index++
				}
				starMap[star] = append(starMap[star], numFromString(l[start:end]))
			}

		}

		lines.prev.forSymbol(func(col int) {
			currentStar := fmt.Sprintf("%d,%d", col, nLine-1)
			numbers := make(map[int64]struct{})
			for i := -1; i < 2; i++ {
				if lines.current.hasNumber(col + i) {
					numbers[lines.current.getNum(col+i)] = struct{}{}
				}
			}
			for num := range numbers {
				starMap[currentStar] = append(starMap[currentStar], num)
			}
		})

		lines.current.forSymbol(func(col int) {
			currentStar := fmt.Sprintf("%d,%d", col, nLine)
			numbers := make(map[int64]struct{})
			for i := -1; i < 2; i++ {
				if lines.prev.hasNumber(col + i) {
					numbers[lines.prev.getNum(col+i)] = struct{}{}
				}
			}
			for num := range numbers {
				starMap[currentStar] = append(starMap[currentStar], num)
			}
		})

		lines.update()

		nLine++

		if err == io.EOF {
			break
		}
	}

	for _, s := range stars {
		for _, star := range s {
			fmt.Println(star, starMap[star])
			if len(starMap[star]) == 2 {
				result += starMap[star][0] * starMap[star][1]
			}
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}

func partOne(reader *bufio.Reader) error {
	var result int64
	result = 0
	var l string
	var start int
	var end int
	lines := newLines()
	nLine := 1

	for {

		lastResult := result

		fmt.Printf("----------------------LINE: %d----------------------'\n", nLine)

		fmt.Println("Current Line: ", lines.current)
		fmt.Println("Previous Line: ", lines.prev)

		chunk, err := reader.ReadBytes('\r')

		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		l = string(chunk)

		fmt.Println(l)
		for i := 0; i < len(l); i++ {

			if l[i] == '.' || l[i] == '\r' || l[i] == '\n' {
				continue
			}

			if util.IsDigit(l[i]) {
				start, end = i, i

				for {
					i++
					if i == len(l) {
						break
					}
					if !util.IsDigit(l[i]) {
						end = i
						i--
						break
					}
				}

				lines.current.addNum(l[start:end], start)

			} else {
				var index int

				if i-1 >= 0 && util.IsDigit(l[i-1]) {
					index = i - 1
					end = index
					for {
						if index == 0 {
							break
						}
						if !util.IsDigit(l[index]) {
							start = index
							break
						}
						index--
					}

					result += numFromString(l[start+1 : end+1])
				}

				if i+1 < len(l) && util.IsDigit(l[i+1]) {
					index = i + 1
					start = index
					for {
						if index == len(l) {
							break
						}
						if !util.IsDigit(l[index]) {
							end = index
							break
						}
						index++
					}
					result += numFromString(l[start:end])
				}

				lines.current.addSymbol(i)

			}

		}

		result += lines.compare()

		lines.update()
		nLine++
		fmt.Println("Result Line: ", result-lastResult)
		fmt.Println("--------------------------------------------")

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}
