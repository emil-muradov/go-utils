package semver_compare

import (
	"fmt"
	"go-utils/internal"
	"strings"
)

type CompareResult string

const (
	GT CompareResult = "GT"
	LT CompareResult = "LT"
	EQ CompareResult = "EQ"
)

func CompareSemvers(left, right string) (CompareResult, error) {
	leftTokens := internal.SliceAtoi(strings.Split(left, "."))
	rightTokens := internal.SliceAtoi(strings.Split(right, "."))
	if len(leftTokens) == 1 {
		return "", fmt.Errorf("invalid semver format. expected: x.y.z, got: %s", left)
	}
	if len(rightTokens) == 1 {
		return "", fmt.Errorf("invalid semver format. expected: x.y.z, got: %s", right)
	}
	n := max(len(leftTokens), len(rightTokens))
	diff := len(leftTokens) - len(rightTokens)
	var tmp []int
	if diff < 0 {
		tmp = make([]int, -diff)
		leftTokens = append(leftTokens, tmp...)
	} else {
		tmp = make([]int, diff)
		rightTokens = append(rightTokens, tmp...)
	}
	for i := range n {
		if leftTokens[i] < rightTokens[i] {
			return LT, nil
		} else if leftTokens[i] > rightTokens[i] {
			return GT, nil
		}
	}
	return EQ, nil
}
