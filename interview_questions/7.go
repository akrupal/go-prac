package main

import "fmt"

func checkExists(e rune, r string) bool {
	for _, el := range r {
		if el == e {
			return true
		}
	}
	return false
}

func removeDuplicates(s string) {
	r := ""
	for _, e := range s {
		if !checkExists(e, r) {
			r = r + string(e)
		}
	}
	fmt.Println(r)
}

func reverseString(s string) {
	r := ""
	for _, e := range s {
		r = string(e) + r
	}
	fmt.Println(r)
}

// func main() {
// 	reverseString("Hello, World!")
// 	removeDuplicates("aabbcc")
// }
