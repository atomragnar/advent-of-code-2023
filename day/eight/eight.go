package eight

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"io"
	"log/slog"
	"regexp"
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func findLCM(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func partTwo(reader *bufio.Reader) error {
	var result int
	NodeCache = make(map[string]*Node)
	var directions []string
	re, _ := regexp.Compile("^\\w\\wA$")
	nodes := make([]*Node, 0)
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
			if re.Match([]byte(node.str)) {
				nodes = append(nodes, node)
			}
			NodeCache[name] = node
		}

		if _, ok := NodeCache[l]; ok {
			ln := NodeCache[l]
			node.left = ln
		} else {
			node.left = NewNode(l)
			if re.Match([]byte(node.left.str)) {
				nodes = append(nodes, node.left)
			}
			NodeCache[l] = node.left
		}

		if _, ok := NodeCache[r]; ok {
			rn := NodeCache[r]
			node.right = rn
		} else {
			node.right = NewNode(r)
			if re.Match([]byte(node.right.str)) {
				nodes = append(nodes, node.right)
			}
			NodeCache[r] = node.right
		}

		if err == io.EOF {
			break
		}

		i++
	}

	var count int
	var keepgoing bool
	re, _ = regexp.Compile("^\\w\\wZ$")
	steps := make([]int, 0)
	for j := 0; j < len(nodes); j++ {
		count = 0
		keepgoing = true
		n := nodes[j]
		for keepgoing {
			i = 0
			for i < len(directions) {
				if directions[i] == "L" {
					n = n.left
				} else {
					n = n.right
				}
				count++
				if re.Match([]byte((*n).str)) {
					fmt.Println(steps)
					steps = append(steps, count)
					keepgoing = false
					break
				}

				i++
			}
		}

	}

	fmt.Println(steps)
	result = findLCM(steps)
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
