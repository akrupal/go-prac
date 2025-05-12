package main

// func main() {
// 	//the point to remember here is if you have a slice with more capacity than it has elements
// 	// the same slice is referenced if you allocate it to some other slice and changing at one place would change at both
// 	arr := make([]int, 0, 5) // Preallocate slice with cap=5
// 	arr = append(arr, 1, 2, 3)
// 	newArr := append(arr, 4) // Capacity is not exceeded
// 	arr[0] = 100
// 	fmt.Println(arr)    // [100 2 3]
// 	fmt.Println(newArr) // [100 2 3 4]
// }
