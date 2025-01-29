package main

import "fmt"

func write(num int) {
	arr := ""

	m := make(map[string]int)
	m["hundred"] = 100
	m["thousands"] = 1000
	m["lakhs"] = 100000
	m["crores"] = 10000000

	starr := []string{"crores", "lakhs", "thousands", "hundred"}
	for _, v := range starr {
		// fmt.Println(num)
		if num == 0 {
			break
		}
		n := num / m[v]
		arr = fmt.Sprintf("%s %d %s", arr, n, v)
		num = num % m[v]
	}
	if num != 0 {
		arr = fmt.Sprintf("%s and %d", arr, num)
	}
	fmt.Println(arr)
}
// func main() {
// 	// 563426789 // 56 crores 34 lakhs 26 thousands 7 hundred and 89
// 	// 21563426789 // 2 thousand 1 hundred and 56 crores 34 lakhs 26 thousands 7 hundred and 89
// 	// 56 crores 34 lakhs 26 thousands 789
// 	// number needs to be printed in words according to the denominations given
// 	write(563426789)
// }
