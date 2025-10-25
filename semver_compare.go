package semver_compare

import (
	"errors"
	"strconv"
	"strings"
)

type CompareResult string

const (
	GT CompareResult = "GT"
	LT CompareResult = "LT"
	EQ CompareResult = "EQ"
)

var (
	ErrInvalidSemver = errors.New("invalid semver. expected x.y.z")
)

func CompareSemvers(left, right string) (CompareResult, error) {
	if left == "" || right == "" {
		return "", ErrInvalidSemver
	}
	left = strings.ReplaceAll(left, ".", "")
	right = strings.ReplaceAll(right, ".", "")
	for len(left) < len(right) {
		left = left + "0"
	}
	for len(left) > len(right) {
		right = right + "0"
	}
	l, err := strconv.Atoi(left)
	if err != nil {
		return "", ErrInvalidSemver
	}
	r, err := strconv.Atoi(right)
	if err != nil {
		return "", ErrInvalidSemver
	}
	if l < r {
		return LT, nil
	}
	if l > r {
		return GT, nil
	}
	return EQ, nil
}
