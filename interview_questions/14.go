package main

//GlobalLogic

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	a := arr[:3]
// 	b := arr[2:5]
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	b := a[:1]
// 	b = append(b, 4)
// 	fmt.Println("Slice a:", a) //1 4 3
// 	fmt.Println("Slice b:", b) //1 4
// }

// func main() {
// 	i := 0
// 	for j := 0; j < 5; j++ {
// 		defer fmt.Println(j)
// 		i++
// 	}
// 	fmt.Println(i)
// 	// 5 4 3 2 1 0
// }

// func main() {
// 	i := 1
// 	defer fmt.Println(i + 1)
// 	i++
// 	fmt.Println("Hello")
// 	//Hello
// 	//2
// 	//2 because its being passed by value to print
// }

// func main() {
// 	ch := make(chan int, 1)//try changing this to 2 or making it unbuffered
// 	ch <- 42
// 	close(ch)

// 	val, ok := <-ch
// 	fmt.Println(val, ok)

// 	val, ok = <-ch
// 	fmt.Println(val, ok)
// }

// difference between primary key unique and foriegn key
// can unique be null?
// difference between where and having in sql
