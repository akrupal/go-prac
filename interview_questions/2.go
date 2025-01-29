package main

import "fmt"

func InsertAtCorrectPlace() {
	strs := []string{"ANIMAL_Cow", "ANIMAL_Goat", "ANIMAL_LION", "BIRD_Crow", "BIRD_Owl", "HUMAN_Aditi", "HUMAN_Ankit", "OTHER_Desk", "OTHER_Pen"}
	inputStr := "ANIMAL_Dog"
	var v int
	for i, s := range strs {
		if s > inputStr {
			v = i
			break
		}
	}
	strs = append(strs[:v], append([]string{inputStr}, strs[v:]...)...)
	fmt.Println(strs)
}

// func main() {
// 	InsertAtCorrectPlace()
// }
