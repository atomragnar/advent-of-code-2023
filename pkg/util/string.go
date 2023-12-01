package util

import (
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
