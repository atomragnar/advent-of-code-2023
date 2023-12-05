package four

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
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
		wins = 0

		util.NumIter[int64](numbers, formatDrawnNumString, util.NumConversion64, func(n int64) {
			if tickets.Contains(int8(n)) {
				wins++
				if _, ok := cardCopies[nLine+wins]; !ok {
					cardCopies[nLine+wins] = 1 + cardCopies[nLine]
				} else {
					cardCopies[nLine+wins] = cardCopies[nLine+wins] + cardCopies[nLine]
				}
			}
		})

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

func formatTicketNumString(s string) string {
	return strings.TrimSpace(util.Split(s, "|")[0])
}

func formatDrawnNumString(s string) string {
	return strings.TrimSpace(util.Split(s, "|")[1])
}

func getTicketNums(s string) *util.Set[int8] {
	set := util.NewSet[int8]()
	util.NumIter[int8](s, formatTicketNumString, util.NumConversion8, func(n int8) {
		set.Add(n)
	})
	return set
}

func partOne(reader *bufio.Reader) error {
	var result int64
	result = 0
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

		lineRes = 0

		util.NumIter[int](numbers, formatDrawnNumString, util.IntConversion, func(n int) {
			if tickets.Contains(int8(n)) {
				if lineRes == 0 {
					lineRes = 1
				} else {
					lineRes += lineRes
				}
			}

		})

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
