package main

import "fmt"

func nextIpAddress(ar []int) {
	for i := 3; i >= 0; i-- {
		if ar[i] < 255 {
			ar[i]++
			break
		} else {
			ar[i] = 0
		}
	}
	fmt.Println(ar)
}

// func main() {
// 	//each value in the ip can vary from 0 to 255
// 	ip := []int{1, 2, 3, 255}
// 	nextIpAddress(ip)
// }
