package args2

import "regexp"

type Args struct {
	lines []string
}

func (a Args) Result() ([][]string) {

	return [][]string{{"l", ""}, {"d", "/opt"}, {"p", "8080"}}
}

func (a Args) len() int {
	return len(a.lines)
}

func (a Args) isValid() (bool) {
	position := 0
	for position < a.len() {
		if match, _ := regexp.MatchString(`-\w`, a.lines[position]); match {

		}
	}
	return true
}

func (a Args) parse() (bool) {
	return true
}
