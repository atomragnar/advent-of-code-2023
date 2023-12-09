package two

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

type color int

const (
	blue color = iota
	red
	green
)

func stringToColor(s string) color {
	switch s {
	case "blue":
		return blue
	case "red":
		return red
	case "green":
		return green
	default:
		return -1
	}
}

func Solution(isPartTwo bool) {
	var fn util.BufferProcessor
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	err := util.ProcessInput(util.DataPath(2), fn)
	if err != nil {
		return
	}
}

func getGameValue(s string) (int64, color) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, " ")
	num, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		slog.Error("Error parsing number", "error", err)
		return 0, -1
	}
	currentColor := stringToColor(split[1])
	return num, currentColor
}

func partOneGame(chunk []byte, constraint map[color]int64) int64 {
	line := strings.Split(string(chunk), ":")
	for _, v := range strings.Split(line[1], ";") {
		v = strings.TrimSpace(v)
		for _, c := range strings.Split(v, ",") {
			num, currentColor := getGameValue(c)
			if num > constraint[currentColor] {
				return 0
			}
		}
	}
	gameNum, err := strconv.ParseInt(line[0][5:], 10, 64)
	if err != nil {
		slog.Error("Error parsing number", "error", err)
		return 0
	}
	fmt.Println(gameNum)
	return gameNum
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
		result += partTwoGame(chunk)
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}

func partTwoGame(chunk []byte) int64 {
	line := strings.Split(string(chunk), ":")
	game := map[color]int64{
		blue:  0,
		red:   0,
		green: 0,
	}
	for _, v := range strings.Split(line[1], ";") {
		v = strings.TrimSpace(v)
		for _, c := range strings.Split(v, ",") {
			num, currentColor := getGameValue(c)
			if num > game[currentColor] {
				game[currentColor] = num
			}
		}
	}
	return game[blue] * game[red] * game[green]
}

func partOne(reader *bufio.Reader) error {
	var constraint = map[color]int64{
		blue:  14,
		red:   12,
		green: 13,
	}
	var result int64
	result = 0

	for {

		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}
		result += partOneGame(chunk, constraint)
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}
