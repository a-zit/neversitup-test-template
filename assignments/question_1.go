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

	permute = func(r []rune, permPos, maxLen int) {
		if permPos == maxLen {
			result = append(result, string(r))
		} else {
			used := make(map[rune]bool)
			for i := permPos; i <= maxLen; i++ {
				if _, ok := used[r[i]]; ok {
					continue
				}
				used[r[i]] = true
				swap(&r, permPos, i)
				permute(r, permPos+1, maxLen)
				swap(&r, permPos, i)
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
