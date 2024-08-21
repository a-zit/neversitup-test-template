package assignments

import (
	"regexp"
)

// countSmileys([':)', ';(', ';}', ':-D']);       // should return 2;
// countSmileys([';D', ':-(', ':-)', ';~)']);     // should return 3;
// countSmileys([';]', ':[', ';*', ':$', ';-D']); // should return 1;

func QuestionThree() {
	countSmileys([]string{":)", ";(", ";}", ":-D"})
	countSmileys([]string{";D", ":-(", ":-)", ";~)"})
	countSmileys([]string{";]", ":[", ";*", ":$", ";-D"})
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

	// fmt.Println(count)
	return count
}
