package assignments

import (
	"fmt"
	"regexp"
)

// countSmileys([':)', ';(', ';}', ':-D']);       // should return 2;
// countSmileys([';D', ':-(', ':-)', ';~)']);     // should return 3;
// countSmileys([';]', ':[', ';*', ':$', ';-D']); // should return 1;

func QuestionThree() {
	fmt.Println("Question 3")
	countSmileys([]string{":)", ";(", ";}", ":-D"})
	countSmileys([]string{";D", ":-(", ":-)", ";~)"})
	countSmileys([]string{";]", ":[", ";*", ":$", ";-D"})
	fmt.Println()
}

func countSmileys(s []string) int {

	regexFormat := `^[:;][-~]?[)D]$`
	regex := regexp.MustCompile(regexFormat)

	count := 0
	for _, v := range s {
		if regex.MatchString(v) {
			count++
		}
	}

	fmt.Println("input:", s, ",result:", count)
	return count
}
