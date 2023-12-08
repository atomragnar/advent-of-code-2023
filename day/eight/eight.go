package eight

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/pkg/util"
	"io"
	"log/slog"
	"strings"
)

var NodeCache map[string]*Node

func Solution(isPartTwo bool) {
	day := 8
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

type Node struct {
	str   string
	left  *Node
	right *Node
}

func NewNode(s string) *Node {
	return &Node{
		str: s,
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
	var result int
	NodeCache = make(map[string]*Node)
	var directions []string
	i := 0

	for {

		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {

			slog.Error("Error reading file", "error", err)
			return err

		}

		if i == 0 {
			directions = strings.Split(strings.TrimSpace(string(chunk)), "")
			i++
			continue
		}

		if len(chunk) <= 2 && (bytes.Contains(chunk, []byte("\r")) || bytes.Contains(chunk, []byte("\n"))) {
			i++
			continue
		}

		sArr := strings.Split(string(chunk), "=")

		name := strings.TrimSpace(sArr[0])

		sArr = strings.Split(sArr[1], ",")

		l := strings.TrimSpace(sArr[0])[1:]

		r := strings.TrimSpace(sArr[1])[:len(sArr[0])-2]
		var node *Node

		if _, ok := NodeCache[name]; ok {
			node = NodeCache[name]
		} else {
			node = NewNode(name)
			NodeCache[name] = node
		}

		if _, ok := NodeCache[l]; ok {
			ln := NodeCache[l]
			node.left = ln
		} else {
			node.left = NewNode(l)
			NodeCache[l] = node.left
		}

		if _, ok := NodeCache[r]; ok {
			rn := NodeCache[r]
			node.right = rn
		} else {
			node.right = NewNode(r)
			NodeCache[r] = node.right
		}

		if err == io.EOF {
			break
		}

		i++
	}

	node := NodeCache["AAA"]
	count := 0
	keepgoing := true
	for keepgoing {

		i = 0

		for i < len(directions) {
			n := *node
			if directions[i] == "L" {
				node = n.left
			} else {
				node = n.right
			}
			count++
			if (*node).str == "ZZZ" {
				fmt.Println(n)
				result = count
				keepgoing = false
				break
			}

			i++
		}

	}

	fmt.Printf("Results is: %d", result)
	return nil
}
