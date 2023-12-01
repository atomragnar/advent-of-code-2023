package one

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
	"strconv"
)

func Solution() {
	err := util.ProcessInput(util.DataPath(1), code)
	if err != nil {
		return
	}
}

func code(reader *bufio.Reader) error {
	var result int64
	result = 0

	for {

		chunk, err := reader.ReadBytes('\n')

		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}

		num := ""
		firstNum := ""
		secondNum := ""
		var length int
		var begin uint8
		var end uint8
		var back int

		line := string(chunk)
		length = len(line)
		back = length - 1
		for i := 0; i < length; i++ {

			if firstNum == "" {
				begin = line[i]
				nBegin := util.ContainsNum(string(line[:i+3]))
				if util.IsDigit(begin) {
					firstNum = string(begin)
				}
				if nBegin != -1 {
					firstNum = strconv.FormatInt(nBegin, 10)
				}
			}

			if secondNum == "" {
				end = line[(back)-i]
				nEnd := util.ContainsNum(line[(back)-i:])
				if util.IsDigit(end) {
					secondNum = string(end)
				}
				if nEnd != -1 {
					secondNum = strconv.FormatInt(nEnd, 10)
				}
			}

			if firstNum != "" && secondNum != "" {
				num = fmt.Sprintf("%s%s", firstNum, secondNum)
				break
			}

		}

		if num != "" {
			numInt, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				return err
			}
			result += numInt
		}

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Results is: %d", result)
	return nil
}
