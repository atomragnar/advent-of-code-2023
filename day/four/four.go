package four

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

func Solution(isPartTwo bool) {
	day := 4
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

func partTwo(reader *bufio.Reader) error {
	var result int64
	result = 0
	var lastRes int64
	var wins int
	nLine := 1
	cardCopies := make(map[int]int64)
	for {

		fmt.Printf("==================LINE: %d =================\n", nLine)
		lastRes = result

		chunk, err := reader.ReadBytes('\n')

		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		if _, ok := cardCopies[nLine]; !ok {
			cardCopies[nLine] = 1
		}

		numbers := util.ByteSplit(chunk, ":")[1]
		tickets := getTicketNums(numbers)
		numbers = strings.TrimSpace(util.Split(numbers, "|")[1])
		wins = 0

		var start int
		i := 0
		for {

			if i >= len(numbers) {
				break
			}

			if util.IsDigit(numbers[i]) {
				start = i
				for {
					if i == len(numbers) || !util.IsDigit(numbers[i]) {
						if num, err := strconv.ParseInt(numbers[start:i], 10, 64); err == nil {
							if tickets.Contains(int8(num)) {
								wins++
								if _, ok := cardCopies[nLine+wins]; !ok {
									cardCopies[nLine+wins] = 1 + cardCopies[nLine]
								} else {
									cardCopies[nLine+wins] = cardCopies[nLine+wins] + cardCopies[nLine]
								}
							}
						} else {
							slog.Error("Error reading file", "error", err)
						}
						break
					}
					i++
				}
			}

			i++

		}

		result += cardCopies[nLine]

		delete(cardCopies, nLine)

		fmt.Printf("Last lines result: %d\n", lastRes)
		fmt.Printf("This lines result: %d\n", result-lastRes)
		fmt.Printf("Total result: %d\n", result)
		fmt.Println("----------------------------------------------")

		if err == io.EOF {
			break
		}

		nLine++
	}

	fmt.Printf("Results is: %d", result)
	return nil
}

func getTicketNums(s string) *util.Set[int8] {
	numString := strings.TrimSpace(util.Split(s, "|")[0])
	set := util.NewSet[int8]()
	var start int
	i := 0
	for {

		if i >= len(numString) {
			break
		}

		if util.IsDigit(numString[i]) {
			start = i
			for {
				if i == len(numString) || !util.IsDigit(numString[i]) {
					if num, err := strconv.ParseInt(numString[start:i], 10, 8); err == nil {
						set.Add(int8(num))
					} else {
						slog.Error("Error reading file", "error", err)
					}
					break
				}
				i++
			}
		}

		i++

	}
	return set
}

func partOne(reader *bufio.Reader) error {
	var result int64
	result = 0
	var start int
	var i int
	var lineRes int
	var lastRes int64
	nLine := 1
	for {

		fmt.Printf("==================LINE: %d =================\n", nLine)
		lastRes = result

		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}

		numbers := util.ByteSplit(chunk, ":")[1]
		tickets := getTicketNums(numbers)

		numbers = strings.TrimSpace(util.Split(numbers, "|")[1])
		lineRes = 0
		i = 0

		for {

			if i >= len(numbers) {
				break
			}

			if util.IsDigit(numbers[i]) {
				start = i
				for {
					if i == len(numbers) || !util.IsDigit(numbers[i]) {
						if num, err := strconv.ParseInt(numbers[start:i], 10, 8); err == nil {
							if tickets.Contains(int8(num)) {
								if lineRes == 0 {
									lineRes = 1
								} else {
									lineRes += lineRes
								}
							}
						} else {
							slog.Error("Error reading file", "error", err)
						}
						break
					}
					i++
				}
			}

			i++

		}

		if err == io.EOF {
			break
		}

		result += int64(lineRes)
		nLine++
		fmt.Printf("Last lines result: %d\n", lastRes)
		fmt.Printf("This lines result: %d\n", lineRes)
		fmt.Printf("Total result: %d\n", result)
		fmt.Println("----------------------------------------------")

	}

	fmt.Printf("Results is: %d", result)
	return nil
}
