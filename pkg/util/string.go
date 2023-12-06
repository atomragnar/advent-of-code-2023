package util

import (
	"log"
	"strconv"
	"strings"
)

func ContainsNum(s string) int64 {
	switch {
	case strings.Contains(s, "one"):
		return 1
	case strings.Contains(s, "two"):
		return 2
	case strings.Contains(s, "three"):
		return 3
	case strings.Contains(s, "four"):
		return 4
	case strings.Contains(s, "five"):
		return 5
	case strings.Contains(s, "six"):
		return 6
	case strings.Contains(s, "seven"):
		return 7
	case strings.Contains(s, "eight"):
		return 8
	case strings.Contains(s, "nine"):
		return 9
	case strings.Contains(s, "zero"):
		return 0
	default:
		return -1
	}
}

func IsDigit(s uint8) bool {
	return s >= '0' && s <= '9'
}

func ByteSplit(b []byte, sep string) []string {
	return strings.Split(strings.TrimSpace(string(b)), sep)
}

func Split(s, sep string) []string {
	return strings.Split(strings.TrimSpace(s), sep)
}

func StrSplitIter(s string, sep string, action func(string)) {
	for _, str := range strings.Split(s, sep) {
		action(str)
	}
}

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// NumIter is a generic function with a type parameter T constrained by Num.
func NumIter[T Num](s string, stringToNums func(s string) string, stringToT func(s string) (T, error), action func(n T)) {
	s = stringToNums(s)
	var start int
	i := 0
	for {
		if i >= len(s) {
			break
		}

		if IsDigit(s[i]) {
			start = i
			for {
				if i == len(s) || !IsDigit(s[i]) {
					numStr := s[start:i]
					if num, err := stringToT(numStr); err == nil {
						action(num)
					} else {
						log.Println("Error converting number:", err)
					}
					break
				}
				i++
			}
		}
		i++
	}
}

func NumConversion8(s string) (int8, error) {
	if num, err := strconv.ParseInt(s, 10, 8); err == nil {
		return int8(num), nil
	} else {
		return 0, err
	}
}

func NumConversion64(s string) (int64, error) {
	if num, err := strconv.ParseInt(s, 10, 8); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func IntConversion(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToInts(s string) []int {
	ints := make([]int, 0)
	stringf := func(z string) string {
		return z
	}
	NumIter[int](s, stringf, IntConversion, func(n int) {
		ints = append(ints, n)
	})
	return ints
}
