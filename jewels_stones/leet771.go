package jewels_stones

import (
	"strings"
)

var (
	j1 = "aA"
	s1 = "aAAbbbb"
)

func Solved() {
	numJewelsInStones(j1, s1)
}

func createMap(jewels []string) map[string]int {
	m := make(map[string]int, len(jewels))
	for _, v := range jewels {
		m[v] = 0
	}
	return m
}

func numJewelsInStones(j string, s string) int {
	var res int
	jewels := strings.Split(j, "")
	m := createMap(jewels)

	for _, v := range s {
		if _, key := m[string(v)]; key {
			res++
		}
	}

	return res
}
