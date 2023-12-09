package five

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"io"
	"log"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

// messy brute force

func Solution(isPartTwo bool) {
	day := 5
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

type seed struct {
	dst int
	r   int
}

type mapper struct {
	src int
	dst int
	r   int
}

func partTwo(reader *bufio.Reader) error {
	r, _ := regexp.Compile("(?:^\\d+\\s\\d+\\s\\d+)")
	seedRegex, _ := regexp.Compile("(?:^seeds:\\s[\\d+\\s]+$)")
	seeds := make([]seed, 0)
	m := make(map[string]map[string][]mapper)
	from := ""
	to := ""

	for {
		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		if seedRegex.Match(chunk) {
			b := strings.Split(string(bytes.TrimSpace(bytes.Split(chunk, []byte(":"))[1])), " ")
			fmt.Println(b)
			for i := 0; i < len(b); i += 2 {
				s := seed{}
				s.dst, err = strconv.Atoi(b[i])
				if err != nil {
					log.Fatal(err)
				}
				s.r, err = strconv.Atoi(strings.ReplaceAll(b[i+1], "\\n\\r", ""))
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, s)
			}
			continue
		}

		if len(chunk) <= 2 && (bytes.Contains(chunk, []byte("\r")) || bytes.Contains(chunk, []byte("\n"))) {
			continue
		}

		if r.Match(chunk) {
			node := mapper{}
			i := 0

			util.StrSplitIter(string(chunk), " ", func(s string) {

				switch i {
				case 0:
					node.dst, err = strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
				case 1:
					node.src, err = strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
				case 2:
					node.r, err = strconv.Atoi(strings.TrimSpace(s))
					if err != nil {
						fmt.Println("is this error")
						log.Fatal(err)
					}
				}
				i++
			})

			m[from][to] = append(m[from][to], node)

		} else if err != io.EOF {

			mapping := bytes.Split(bytes.Split(chunk, []byte(" "))[0], []byte("-"))
			from = string(mapping[0])
			to = string(mapping[2])

			if _, ok := m[from]; !ok {
				m[from] = make(map[string][]mapper)
				m[from][to] = make([]mapper, 0)

			} else if _, ok := m[from][to]; !ok {
				m[from][to] = make([]mapper, 0)

			} else {
				// TODO - error
			}
		}

		if err == io.EOF {
			break
		}
	}

	loc := 0

	for {

		h := getValue(m["humidity"]["location"], loc)
		t := getValue(m["temperature"]["humidity"], h)
		l := getValue(m["light"]["temperature"], t)
		w := getValue(m["water"]["light"], l)
		f := getValue(m["fertilizer"]["water"], w)
		soil := getValue(m["soil"]["fertilizer"], f)
		s := getValue(m["seed"]["soil"], soil)

		for _, seed := range seeds {
			if s >= seed.dst && s < seed.r+seed.dst {
				fmt.Println(loc)
				return nil
			}

		}
		loc += 1
	}

	return nil
}

func getValue(slice []mapper, n int) int {
	for _, m := range slice {
		if n >= m.dst && n < m.dst+m.r {
			return m.src + (n - m.dst)
		}
	}

	return n
}

func (m *mapper) intRange(num int) bool {
	if num < m.src {
		return false
	}
	if num > m.src+m.r {
		return false
	}
	if num == m.src+m.r {
		return false
	}
	return true
}

func (m *mapper) getIntVal(val int) int {
	n := val - m.src
	n = n + m.dst
	return n
}

func partOne(reader *bufio.Reader) error {
	r, _ := regexp.Compile("(?:^\\d+\\s\\d+\\s\\d+)")
	seedRegex, _ := regexp.Compile("(?:^seeds:\\s[\\d+\\s]+$)")
	seeds := make([]mapper, 0)
	m := make(map[string]map[string][]mapper)
	from := ""
	to := ""
	for {
		chunk, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			slog.Error("Error reading file", "error", err)
			return err
		}

		if seedRegex.Match(chunk) {
			b := bytes.TrimSpace(bytes.Split(chunk, []byte(":"))[1])
			util.StrSplitIter(string(b), " ", func(s string) {
				seed := mapper{}
				seed.src, err = strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, seed)
			})
			continue
		}

		if len(chunk) <= 2 && (bytes.Contains(chunk, []byte("\r")) || bytes.Contains(chunk, []byte("\n"))) {
			continue
		}

		if r.Match(chunk) {
			node := mapper{}
			i := 0

			util.StrSplitIter(string(chunk), " ", func(s string) {

				switch i {
				case 0:
					node.dst, err = strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
				case 1:
					node.src, err = strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
				case 2:
					node.r, err = strconv.Atoi(strings.TrimSpace(s))
					if err != nil {
						fmt.Println("is this error")
						log.Fatal(err)
					}
				}
				i++
			})

			m[from][to] = append(m[from][to], node)

		} else if err != io.EOF {

			mapping := bytes.Split(bytes.Split(chunk, []byte(" "))[0], []byte("-"))
			from = string(mapping[0])
			to = string(mapping[2])

			if _, ok := m[from]; !ok {
				m[from] = make(map[string][]mapper)
				m[from][to] = make([]mapper, 0)

			} else if _, ok := m[from][to]; !ok {
				m[from][to] = make([]mapper, 0)

			} else {
				// TODO - error
			}
		}

		if err == io.EOF {
			break
		}
	}

	low := -1
	for _, seed := range seeds {
		currentType := "seed"
		currentValue := seed.src
		for {
			currMap := m[currentType]
			isSet := false
			for to, nodes := range currMap {
				if !isSet {
					currentType = to
					isSet = true
				}
				for _, node := range nodes {
					if node.intRange(currentValue) {
						currentValue = node.getIntVal(currentValue)
						goto NEXT
					}
				}
			}
		NEXT:
			if currentType == "location" {
				if low == -1 {
					low = currentValue
				}
				if currentValue < low {
					low = currentValue
				}
				break
			}

		}

	}
	fmt.Printf("Results is: %d", low)
	return nil
}
