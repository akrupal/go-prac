package main

// import "fmt"

func retNum(ascendingChan chan int, descendingChan chan int, targetSum int, chanSize int) int {
	ascending := make([]int, chanSize)
	descending := make([]int, chanSize)

	for i := 0; i < chanSize; i++ {
		ascending[i] = <-ascendingChan
	}

	for i := 0; i < chanSize; i++ {
		descending[i] = <-descendingChan
	}

	i := 0
	j := 0
	count := 0

	for i < chanSize && j < chanSize {
		sum := ascending[i] + descending[j]

		if sum == targetSum {
			count++
			i++
			j++
		} else if sum < targetSum {
			i++
		} else {
			j++
		}
	}

	return count
}

func findArraySum(arr []int32) int32 {
	const mod int64 = 1000000007

	var sum int64

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum += int64(arr[j] - arr[i])
		}
	}

	sum %= mod

	// Handle negative modulo
	if sum < 0 {
		sum += mod
	}

	return int32(sum)
}

// func main() {
// 	chan1 := make(chan int, 5)
// 	chan2 := make(chan int, 5)

// 	chan1 <- 0
// 	chan1 <- 2
// 	chan1 <- 3
// 	chan1 <- 5
// 	chan1 <- 7

// 	chan2 <- 9
// 	chan2 <- 8
// 	chan2 <- 6
// 	chan2 <- 5
// 	chan2 <- 1

// 	fmt.Println(retNum(chan1, chan2, 10, 5))

// 	arr1 := []int32{1, 2, 3}
// 	fmt.Println(findArraySum(arr1))

// 	arr := []int32{4, 1, 3, 2}
// 	fmt.Println(findArraySum(arr))
// }
