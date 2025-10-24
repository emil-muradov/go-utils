package internal

import "strconv"

func SliceAtoi(s []string) []int {
	res := make([]int, len(s))
	for i := range len(res) {
		v, _ := strconv.Atoi(s[i])
		res[i] = v
	}
	return res
}