package assignments

import "fmt"

func QuestionOne() {
	Permutations("aabb")
	// Permutations("ab")
	// Permutations("abc")
	// Permutations("a")
}

func Permutations(s string) []string {
	var result []string
	chars := []rune(s)
	var permute func([]rune, int, int)
	existed := make(map[string]bool)

	permute = func(r []rune, left, right int) {
		if left == right {
			if !existed[string(r)] {
				// fmt.Println("take:", string(r))
				result = append(result, string(r))
				existed[string(r)] = true
			}
		} else {
			for i := left; i <= right; i++ {
				swap(&r, left, i)
				// fmt.Println("left-i:", left, i, ", result:", string(r))
				permute(r, left+1, right)
				swap(&r, left, i)
				// fmt.Println("result-2:", string(r))
			}
		}
	}

	permute(chars, 0, len(chars)-1)

	fmt.Println(result)
	return result
}

func swap(s *[]rune, i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
