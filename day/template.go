package day

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
)

func Solution(isPartTwo bool) {
	day := 0
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

	for {

		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}
		fmt.Println(string(chunk))
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}

func partOne(reader *bufio.Reader) error {
	var result int64
	result = 0

	for {

		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}
		fmt.Println(string(chunk))
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}
