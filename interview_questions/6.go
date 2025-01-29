package main

import (
	"fmt"
	"sort"
)

func sortedString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func coupleSimilarStrings() {
	arr := []string{"tac", "act", "cat", "goc", "cog", "ogc"}
	m := make(map[string][]string)
	var g [][]string

	for _, a := range arr {
		key := sortedString(a)
		m[key] = append(m[key], a)
	}

	for _, v := range m {
		g = append(g, v)
	}
	fmt.Println(g)
}

// func main() {
// 	coupleSimilarStrings()
// }
