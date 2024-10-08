package assignments

import "fmt"

// [7] should return 7, because it occurs 1 time (which is odd).
// [0] should return 0, because it occurs 1 time (which is odd).
// [1,1,2] should return 2, because it occurs 1 time (which is odd).
// [0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
// [1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).

func QuestionTwo() {
	fmt.Println("Question 2")
	FindOdd([]int{7})
	FindOdd([]int{0})
	FindOdd([]int{1, 1, 2})
	FindOdd([]int{0, 1, 0, 1, 0})
	FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
	fmt.Println()
}

func FindOdd(arr []int) int {
	result := 0
	dict := make(map[int]int)

	for _, v := range arr {
		dict[v]++
	}

	for k, v := range dict {
		if v%2 != 0 {
			result = k
			break
		}
	}

	fmt.Println("input:", arr, ",result:", result)
	return result
}
