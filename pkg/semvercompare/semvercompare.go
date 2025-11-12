package semvercompare

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CompareResult string

const (
	GT CompareResult = "GT"
	LT CompareResult = "LT"
	EQ CompareResult = "EQ"
)

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
const semverPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
const digitPattern = `\d+`

var (
	ErrInvalidSemver = errors.New("invalid semver. expected x.y.z")
)

func CompareSemvers(left, right string) (CompareResult, error) {
	semverRegexp, err := regexp.Compile(semverPattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile semver pattern: %w", err)
	}
	if !semverRegexp.MatchString(left) || !semverRegexp.MatchString(right) {
		return "", ErrInvalidSemver
	}
	digitRegexp, err := regexp.Compile(digitPattern)
	if err != nil {
		return "", fmt.Errorf("failed to compile digit pattern: %w", err)
	}
	leftDigits := strings.Join(digitRegexp.FindAllString(left, -1), "")
	rightDigits := strings.Join(digitRegexp.FindAllString(right, -1), "")
	for len(leftDigits) < len(rightDigits) {
		leftDigits = leftDigits + "0"
	}
	for len(leftDigits) > len(rightDigits) {
		rightDigits = rightDigits + "0"
	}
	l, _ := strconv.Atoi(leftDigits)
	r, _ := strconv.Atoi(rightDigits)
	if l < r {
		return LT, nil
	}
	if l > r {
		return GT, nil
	}
	return EQ, nil
}
