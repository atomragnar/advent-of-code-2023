package day

//
//import (
//	"fmt"
//	"github.com/atomragnar/advent-of-code-2023/pkg/util"
//	"log/slog"
//	"os"
//)
//
//func Solution(isPartTwo bool) {
//	day := 0
//	data := func() ([]byte, error) {
//		d, err := os.ReadFile(util.DataPath(day))
//		if err != nil {
//			return nil, err
//		}
//		return d, err
//	}
//	var fn func(func() ([]byte, error)) error
//	if isPartTwo {
//		fn = partTwo
//	} else {
//		fn = partOne
//	}
//	err := fn(data)
//	if err != nil {
//		return
//	}
//}
//
//func partTwo(input func() ([]byte, error)) error {
//	var result int64
//	result = 0
//
//	data, err := input()
//	if err != nil {
//		slog.Error("Error reading file", "error", err)
//		return err
//	}
//
//	fmt.Println(string(data))
//
//	fmt.Printf("Results is: %d", result)
//	return nil
//}
//
//func partOne(input func() ([]byte, error)) error {
//	var result int64
//	result = 0
//
//	data, err := input()
//	if err != nil {
//		slog.Error("Error reading file", "error", err)
//		return err
//	}
//
//	fmt.Println(string(data))
//
//	fmt.Printf("Results is: %d", result)
//	return nil
//}
