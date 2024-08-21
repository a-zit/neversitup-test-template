package assignments

import "fmt"

func QuestionOne() {
	fmt.Println("Question 1")
	Permutations("aabb")
	Permutations("ab")
	Permutations("abc")
	Permutations("a")
	fmt.Println()
}

func Permutations(s string) []string {
	var result []string
	chars := []rune(s)
	var permute func([]rune, int, int)
	existed := make(map[string]bool)

	permute = func(r []rune, left, right int) {
		if left == right {
			if !existed[string(r)] {
				result = append(result, string(r))
				existed[string(r)] = true
			}
		} else {
			for i := left; i <= right; i++ {
				swap(&r, left, i)
				permute(r, left+1, right)
				swap(&r, left, i)
			}
		}
	}

	permute(chars, 0, len(chars)-1)

	fmt.Println("input:", s, ",result:", result)
	return result
}

func swap(s *[]rune, i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
