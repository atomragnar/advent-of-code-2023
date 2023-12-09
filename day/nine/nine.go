package nine

import (
	"fmt"
	"github.com/atomragnar/advent-of-code-2023/day/util"
	"log/slog"
	"os"
)

func Solution(isPartTwo bool) {
	day := 9
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

func extrapolate(nums []int) (int, int) {
	d := make([]int, len(nums)-1)
	z := true
	for i := 0; i < len(nums)-1; i++ {
		d[i] = nums[i+1] - nums[i]
		if d[i] != 0 {
			z = false
		}
	}
	first := 0
	last := 0
	if !z {
		last, first = extrapolate(d)
	}
	return nums[len(nums)-1] + last, nums[0] - first
}

func partTwo(input func() ([]byte, error)) error {
	partOne(input)
	return nil
}

func partOne(input func() ([]byte, error)) error {
	result1, result2 := 0, 0
	var nums []int

	data, err := input()
	if err != nil {
		slog.Error("Error reading file", "error", err)
		return err
	}

	sArr := util.SplitData(data)

	for _, s := range sArr {
		nums = util.StringToInts(s)
		fmt.Println(nums)
		l, f := extrapolate(nums)
		result1 = result1 + l
		result2 = result2 + f
	}

	fmt.Printf("Results is: %d, %d", result1, result2)
	return nil
}
