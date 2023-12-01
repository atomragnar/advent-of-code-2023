package util

import "fmt"

func DataPath(day int) string {
	return fmt.Sprintf(".\\data\\%d.txt", day)
}
