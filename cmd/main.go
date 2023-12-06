package main

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/five"
	"github.com/atomragnar/advent-of-code-2023/day/four"
	"github.com/atomragnar/advent-of-code-2023/day/one"
	"github.com/atomragnar/advent-of-code-2023/day/three"
	"github.com/atomragnar/advent-of-code-2023/day/two"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type adventFunc func(bool)

var adventFuncs = map[int]adventFunc{
	1: one.Solution,
	2: two.Solution,
	3: three.Solution,
	4: four.Solution,
	5: five.Solution,
}

func main() {
	newDay := false
	if newDay {
		addDay(2)
	}
	adventFuncs[5](true)
}

const templateString = `package {{.PackageName}}

import (
	"bufio"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
)

func Solution(isPartTwo bool) {
	day := {{.DayNumber}}
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

`

const template2String = `package {{.PackageName}}

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"log/slog"
	"os"
)

func Solution(isPartTwo bool) {
	day := {{.DayNumber}}
	data := func() ([]byte, error) {
		d, err := os.ReadFile(util.DataPath(day))
		if err != nil {
			return nil, err
		}
		return d, err
	}
	var fn func(func() ([]byte, error)) error
	if isPartTwo {
		fn = partTwo
	} else {
		fn = partOne
	}
	err := fn(data)
	if err != nil {
		return
	}
}

func partTwo(input func() ([]byte, error)) error {
	var result int64
	result = 0

	data, err := input()
	if err != nil {
		slog.Error("Error reading file", "error", err)
		return err
	}

	fmt.Println(string(data))

	fmt.Printf("Results is: %d", result)
	return nil
}

func partOne(input func() ([]byte, error)) error {
	var result int64
	result = 0

	data, err := input()
	if err != nil {
		slog.Error("Error reading file", "error", err)
		return err
	}

	fmt.Println(string(data))

	fmt.Printf("Results is: %d", result)
	return nil
}
`

type TemplateData struct {
	PackageName string
	DayNumber   int
}

func addDay(ver int) {
	//dayName := time.Now().Format("Monday")
	dayNumber := time.Now().Day()
	dayName := numberToText(dayNumber)

	data := TemplateData{
		PackageName: strings.ToLower(dayName),
		DayNumber:   dayNumber,
	}

	var tmplstr string
	if ver == 1 {
		tmplstr = templateString
	} else {
		tmplstr = template2String
	}

	dirPath := filepath.Join(".", "day", data.PackageName)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	filePath := filepath.Join(dirPath, data.PackageName+".go")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	tmpl, err := template.New("goTemplate").Parse(tmplstr)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func numberToText(number int) string {
	numbers := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
		6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten",
		11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen",
		19: "nineteen", 20: "twenty", 21: "twentyone", 22: "twentytwo",
		23: "twentythree", 24: "twentyfour",
		25: "twentyfive",
	}

	if text, ok := numbers[number]; ok {
		return text
	}
	return "unknown"
}
